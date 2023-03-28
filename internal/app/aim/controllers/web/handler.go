package web

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aim_api "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/proto"
)

type Boundary interface {
	bdy.CustomAdder
	bdy.ButtonAdder
	bdy.ToggleAdder
	bdy.ThermostatAdder
	bdy.CommandAdder
	bdy.CustomGetter
	bdy.ToggleGetter
	bdy.ButtonGetter
	bdy.ThermostatGetter
	bdy.AppliancesGetter
	bdy.CommandGetter
	bdy.ApplianceRenamer
	bdy.IRDeviceChanger
	bdy.CommandRenamer
	bdy.IRDataSetter
	bdy.ApplianceDeleter
	bdy.CommandDeleter
}

type Handler struct {
	aim_api.UnimplementedAimServiceServer
	i Boundary
}

var _ aim_api.AimServiceServer = &Handler{}

func NewHandler(i Boundary) *Handler {
	return &Handler{
		i: i,
	}
}

func (h *Handler) AddCustom(ctx context.Context, req *aim_api.AddCustomRequest) (*aim_api.AddAppResponse, error) {
	res := aim_api.AddAppResponse{}
	_res, err := h.i.AddCustom(ctx, bdy.AddCustomInput{
		Name:     req.Name,
		DeviceID: req.DeviceId,
	})

	if err != nil {
		return &res, err
	}
	res.ApplianceId = _res.ID
	return &res, nil
}

func (h *Handler) AddToggle(ctx context.Context, req *aim_api.AddToggleRequest) (*aim_api.AddAppResponse, error) {
	res := aim_api.AddAppResponse{}
	_res, err := h.i.AddToggle(ctx, bdy.AddToggleInput{
		Name:     req.Name,
		DeviceID: req.DeviceId,
	})
	if err != nil {
		return &res, err
	}

	res.ApplianceId = _res.ID
	return &res, nil
}

func (h *Handler) AddButton(ctx context.Context, req *aim_api.AddButtonRequest) (*aim_api.AddAppResponse, error) {
	res := aim_api.AddAppResponse{}
	_res, err := h.i.AddButton(ctx, bdy.AddButtonInput{
		Name:     req.Name,
		DeviceID: req.DeviceId,
	})
	if err != nil {
		return &res, err
	}

	res.ApplianceId = _res.ID
	return &res, nil
}

func (h *Handler) AddThermostat(ctx context.Context, req *aim_api.AddThermostatRequest) (*aim_api.AddAppResponse, error) {
	res := aim_api.AddAppResponse{}
	_res, err := h.i.AddThermostat(ctx, bdy.AddThermostatInput{
		Name:               req.Name,
		DeviceID:           req.DeviceId,
		Scale:              float64(req.TempScale),
		MaximumHeatingTemp: int(req.MaximumHeatingTemp),
		MinimumHeatingTemp: int(req.MinimumHeatingTemp),
		MaximumCoolingTemp: int(req.MaximumCoolingTemp),
		MinimumCoolingTemp: int(req.MinimumCoolingTemp),
	})
	if err != nil {
		return &res, err
	}

	res.ApplianceId = _res.ID
	return &res, nil
}

func (h *Handler) AddCommand(ctx context.Context, req *aim_api.AddCommandRequest) (*empty.Empty, error) {
	res := empty.Empty{}
	err := h.i.AddCommand(ctx, bdy.AddCommandInput{
		AppID: req.ApplianceId,
		Name:  req.Name,
	})
	return &res, err
}

func (h *Handler) GetCustom(ctx context.Context, req *aim_api.GetApplianceRequest) (*aim_api.Custom, error) {
	res := aim_api.Custom{}
	_res, err := h.i.GetCustom(ctx, bdy.GetAppInput{
		AppID: req.ApplianceId,
	})
	if err != nil {
		return &res, err
	}

	res.Name = _res.Name
	res.DeviceId = _res.DeviceID
	res.Id = _res.ID
	res.Commands = convertCommands(_res.Commands)
	return &res, nil
}

func (h *Handler) GetToggle(ctx context.Context, req *aim_api.GetApplianceRequest) (*aim_api.Toggle, error) {
	res := aim_api.Toggle{}

	_res, err := h.i.GetToggle(ctx, bdy.GetAppInput{
		AppID: req.ApplianceId,
	})
	if err != nil {
		return &res, err
	}

	res.Name = _res.Name
	res.DeviceId = _res.DeviceID
	res.Id = _res.ID
	res.Commands = convertCommands(_res.Commands)
	return &res, nil
}

func (h *Handler) GetButton(ctx context.Context, req *aim_api.GetApplianceRequest) (*aim_api.Button, error) {
	res := aim_api.Button{}

	_res, err := h.i.GetButton(ctx, bdy.GetAppInput{
		AppID: req.ApplianceId,
	})
	if err != nil {
		return &res, err
	}

	res.Name = _res.Name
	res.DeviceId = _res.DeviceID
	res.Id = _res.ID
	res.Commands = convertCommands(_res.Commands)
	return &res, nil
}

func (h *Handler) GetThermostat(ctx context.Context, req *aim_api.GetApplianceRequest) (*aim_api.Thermostat, error) {
	res := aim_api.Thermostat{}

	_res, err := h.i.GetThermostat(ctx, bdy.GetAppInput{
		AppID: req.ApplianceId,
	})
	if err != nil {
		return &res, err
	}

	res.Id = _res.ID
	res.Name = _res.Name
	res.DeviceId = _res.DeviceID
	res.TempScale = float32(_res.Scale)
	res.MaximumHeatingTemp = uint32(_res.MaximumHeatingTemp)
	res.MinimumHeatingTemp = uint32(_res.MinimumHeatingTemp)
	res.MaximumCoolingTemp = uint32(_res.MaximumCoolingTemp)
	res.MinimumCoolingTemp = uint32(_res.MinimumCoolingTemp)
	res.Commands = convertCommands(_res.Commands)
	return &res, nil
}

func (h *Handler) GetAppliances(ctx context.Context, _ *empty.Empty) (*aim_api.GetAppliancesResponse, error) {
	res := aim_api.GetAppliancesResponse{}

	_res, err := h.i.GetAppliances(ctx)
	if err != nil {
		return &res, err
	}

	res.Appliances = make([]*aim_api.Appliance, len(_res.Apps))
	for i, a := range _res.Apps {
		res.Appliances[i] = &aim_api.Appliance{
			Id:       a.ID,
			DeviceId: a.DeviceID,
			Type:     convertType(a.ApplianceType),
			Name:     a.Name,
		}
	}
	return &res, nil
}

func (h *Handler) GetCommand(ctx context.Context, req *aim_api.GetCommandRequest) (*aim_api.GetCommandResponse, error) {
	res := aim_api.GetCommandResponse{}

	_res, err := h.i.GetCommand(ctx, bdy.GetCommandInput{
		AppID: req.ApplianceId,
		ComID: req.CommandId,
	})
	if err != nil {
		return &res, err
	}

	res.Name = _res.Name
	res.CommandId = _res.ID
	err = proto.Unmarshal(_res.Data, res.Irdata)
	return &res, err
}

func (h *Handler) RenameAppliance(ctx context.Context, req *aim_api.RenameApplianceRequest) (*empty.Empty, error) {
	res := empty.Empty{}

	err := h.i.RenameAppliance(ctx, bdy.RenameAppInput{
		AppID: req.ApplianceId,
		Name:  req.Name,
	})

	return &res, err
}

func (h *Handler) ChangeDevice(ctx context.Context, req *aim_api.ChangeDeviceRequest) (*empty.Empty, error) {
	res := empty.Empty{}

	err := h.i.ChangeIRDevice(ctx, bdy.ChangeIRDevInput{
		AppID:    req.ApplianceId,
		DeviceID: req.DeviceId,
	})

	return &res, err
}

func (h *Handler) RenameCommand(ctx context.Context, req *aim_api.RenameCommandRequest) (*empty.Empty, error) {
	res := empty.Empty{}

	err := h.i.RenameCommand(ctx, bdy.RenameCommandInput{
		AppID: req.ApplianceId,
		ComID: req.CommandId,
		Name:  req.Name,
	})

	return &res, err
}

func (h *Handler) SetIrData(ctx context.Context, req *aim_api.SetIRDataRequest) (*empty.Empty, error) {
	res := empty.Empty{}

	irdata, err := proto.Marshal(req.Irdata)
	if err != nil {
		return &res, err
	}

	err = h.i.SetIRData(ctx, bdy.SetIRDataInput{
		AppID: req.ApplianceId,
		ComID: req.CommandId,
		Data:  irdata,
	})

	return &res, err
}

func (h *Handler) DeleteAppliance(ctx context.Context, req *aim_api.DeleteApplianceRequest) (*empty.Empty, error) {
	res := empty.Empty{}

	err := h.i.DeleteAppliance(ctx, bdy.DeleteAppInput{
		AppID: req.ApplianceId,
	})

	return &res, err
}

func (h *Handler) DeleteCommand(ctx context.Context, req *aim_api.DeleteCommandRequest) (*empty.Empty, error) {
	res := empty.Empty{}

	err := h.i.DeleteCommand(ctx, bdy.DeleteCommandInput{
		AppID: req.ApplianceId,
		ComID: req.CommandId,
	})

	return &res, err
}

func (h *Handler) NotifyApplianceUpdate(*empty.Empty, aim_api.AimService_NotifyApplianceUpdateServer) error {
	return nil
}

func convertCommands(coms []bdy.Command) []*aim_api.Command {
	res := make([]*aim_api.Command, len(coms))
	for i, c := range coms {
		res[i] = &aim_api.Command{
			Id:   c.ID,
			Name: c.Name,
		}
	}
	return res
}

func convertType(appType bdy.ApplianceType) aim_api.Appliance_ApplianceType {
	switch appType {
	case bdy.TypeCustom:
		return aim_api.Appliance_APPLIANCE_TYPE_CUSTOM
	case bdy.TypeButton:
		return aim_api.Appliance_APPLIANCE_TYPE_BUTTON
	case bdy.TypeToggle:
		return aim_api.Appliance_APPLIANCE_TYPE_TOGGLE
	case bdy.TypeThermostat:
		return aim_api.Appliance_APPLIANCE_TYPE_THERMOSTAT
	}
	return aim_api.Appliance_APPLIANCE_TYPE_UNSPECIFIED
}
