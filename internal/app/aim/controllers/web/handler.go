package web

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aim_api "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/proto"
)

type Handler struct {
	c Controller
	i boundary.InputBoundary
}

func (h *Handler) AddCustom(ctx context.Context, req *aim_api.AddCustomRequest) (*aim_api.AddAppResponse, error) {
	res := aim_api.AddAppResponse{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.AddCustom(ctx, boundary.AddCustomInput{
		Name:     req.Name,
		DeviceID: req.DeviceId,
	})

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	data := _res.Data.(boundary.AddAppOutput)
	res.Id = data.ID
	return &res, _res.Err
}

func (h *Handler) AddToggle(ctx context.Context, req *aim_api.AddToggleRequest) (*aim_api.AddAppResponse, error) {
	res := aim_api.AddAppResponse{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.AddToggle(ctx, boundary.AddToggleInput{
		Name:     req.Name,
		DeviceID: req.DeviceId,
	})

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	data := _res.Data.(boundary.AddAppOutput)
	res.Id = data.ID
	return &res, _res.Err
}

func (h *Handler) AddButton(ctx context.Context, req *aim_api.AddButtonRequest) (*aim_api.AddAppResponse, error) {
	res := aim_api.AddAppResponse{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.AddCustom(ctx, boundary.AddCustomInput{
		Name:     req.Name,
		DeviceID: req.DeviceId,
	})

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	data := _res.Data.(boundary.AddAppOutput)
	res.Id = data.ID
	return &res, _res.Err
}

func (h *Handler) AddThermostat(ctx context.Context, req *aim_api.AddThermostatRequest) (*aim_api.AddAppResponse, error) {
	res := aim_api.AddAppResponse{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.AddThermostat(ctx, boundary.AddThermostatInput{
		Name:               req.Name,
		DeviceID:           req.DeviceId,
		Scale:              float64(req.TempScale),
		MaximumHeatingTemp: int(req.MaximumHeatingTemp),
		MinimumHeatingTemp: int(req.MinimumHeatingTemp),
		MaximumCoolingTemp: int(req.MaximumCoolingTemp),
		MinimumCoolingTemp: int(req.MinimumCoolingTemp),
	})

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	data := _res.Data.(boundary.AddAppOutput)
	res.Id = data.ID
	return &res, _res.Err
}

func (h *Handler) AddCommand(ctx context.Context, req *aim_api.AddCommandRequest) (*empty.Empty, error) {
	res := empty.Empty{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.AddCommand(ctx, boundary.AddCommandInput{
		AppID: req.ApplianceId,
		Name:  req.Name,
	})

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	return &res, _res.Err
}

func (h *Handler) GetCustom(ctx context.Context, req *aim_api.GetApplianceRequest) (*aim_api.Custom, error) {
	res := aim_api.Custom{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.GetCustom(ctx, boundary.GetAppInput{
		AppID: req.Id,
	})

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	data := _res.Data.(boundary.GetCustomOutput)
	res.Name = data.Name
	res.DeviceId = data.DeviceID
	res.Id = data.ID
	//add command converter!!!
	return &res, nil
}

func (h *Handler) GetToggle(ctx context.Context, req *aim_api.GetApplianceRequest) (*aim_api.Toggle, error) {
	res := aim_api.Toggle{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.GetToggle(ctx, boundary.GetAppInput{
		AppID: req.Id,
	})

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	data := _res.Data.(boundary.GetToggleOutput)
	res.Name = data.Name
	res.DeviceId = data.DeviceID
	res.Id = data.ID
	//add command converter!!!
	return &res, nil
}

func (h *Handler) GetButton(ctx context.Context, req *aim_api.GetApplianceRequest) (*aim_api.Button, error) {
	res := aim_api.Button{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.GetButton(ctx, boundary.GetAppInput{
		AppID: req.Id,
	})

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	data := _res.Data.(boundary.GetButtonOutput)
	res.Name = data.Name
	res.DeviceId = data.DeviceID
	res.Id = data.ID
	//add command converter!
	return &res, nil
}

func (h *Handler) GetThermostat(ctx context.Context, req *aim_api.GetApplianceRequest) (*aim_api.Thermostat, error) {
	res := aim_api.Thermostat{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.GetThermostat(ctx, boundary.GetAppInput{
		AppID: req.Id,
	})

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	data := _res.Data.(boundary.GetThermostatOutput)
	res.Id = data.ID
	res.Name = data.Name
	res.DeviceId = data.DeviceID
	res.TempScale = float32(data.Scale)
	res.MaximumHeatingTemp = uint32(data.MaximumHeatingTemp)
	res.MinimumHeatingTemp = uint32(data.MinimumHeatingTemp)
	res.MaximumCoolingTemp = uint32(data.MaximumCoolingTemp)
	res.MinimumCoolingTemp = uint32(data.MinimumCoolingTemp)
	//add command converter!
	return &res, nil
}

func (h *Handler) GetAppliances(ctx context.Context, _ *empty.Empty) (*aim_api.GetAppliancesResponse, error) {
	res := aim_api.GetAppliancesResponse{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.GetAppliances(ctx)

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	data := _res.Data.(boundary.GetAppliancesOutput)
	res.Appliances = make([]*aim_api.Appliance, len(data.Apps))
	for i, a := range data.Apps {
		res.Appliances[i].Id = a.ID
		res.Appliances[i].DeviceId = a.DeviceID
		// add type converter
		res.Appliances[i].Name = a.Name
	}

	return &res, nil
}

func (h *Handler) GetCommand(ctx context.Context, req *aim_api.GetCommandRequest) (*aim_api.GetCommandResponse, error) {
	res := aim_api.GetCommandResponse{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.GetCommand(ctx, boundary.GetCommandInput{
		AppID: req.ApplianceId,
		ComID: req.CommandId,
	})

	_res := (<-ch)
	if _res.Err != nil {
		return &res, _res.Err
	}

	data := _res.Data.(boundary.GetCommandOutput)
	res.Name = data.Name
	res.CommandId = data.ID
	err := proto.Unmarshal(data.Data, res.Irdata)

	return &res, err
}

func (h *Handler) RenameAppliance(ctx context.Context, req *aim_api.RenameApplianceRequest) (*empty.Empty, error) {
	res := empty.Empty{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.RenameAppliance(ctx, boundary.RenameAppInput{
		AppID: req.Id,
		Name:  req.Name,
	})

	_res := (<-ch)
	return &res, _res.Err
}

func (h *Handler) ChangeDevice(ctx context.Context, req *aim_api.ChangeDeviceRequest) (*empty.Empty, error) {
	res := empty.Empty{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.ChangeIRDevice(ctx, boundary.ChangeIRDevInput{
		AppID:    req.Id,
		DeviceID: req.DeviceId,
	})

	_res := (<-ch)
	return &res, _res.Err
}

func (h *Handler) RenameCommand(ctx context.Context, req *aim_api.RenameCommandRequest) (*empty.Empty, error) {
	res := empty.Empty{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.RenameCommand(ctx, boundary.RenameCommandInput{
		AppID: req.ApplianceId,
		ComID: req.CommandId,
		Name:  req.Name,
	})

	_res := (<-ch)
	return &res, _res.Err
}

func (h *Handler) SetIrData(ctx context.Context, req *aim_api.SetIRDataRequest) (*empty.Empty, error) {
	res := empty.Empty{}
	ctx, ch := h.c.NewSession(ctx)

	irdata, err := proto.Marshal(req.Irdata)
	if err != nil {
		return &res, err
	}
	h.i.SetIRData(ctx, boundary.SetIRDataInput{
		AppID: req.ApplianceId,
		ComID: req.CommandId,
		Data:  irdata,
	})

	_res := (<-ch)
	return &res, _res.Err
}

func (h *Handler) DeleteAppliance(ctx context.Context, req *aim_api.DeleteApplianceRequest) (*empty.Empty, error) {
	res := empty.Empty{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.DeleteAppliance(ctx, boundary.DeleteAppInput{
		AppID: req.ApplianceId,
	})

	_res := (<-ch)
	return &res, _res.Err
}

func (h *Handler) DeleteCommand(ctx context.Context, req *aim_api.DeleteCommandRequest) (*empty.Empty, error) {
	res := empty.Empty{}
	ctx, ch := h.c.NewSession(ctx)
	h.i.DeleteCommand(ctx, boundary.DeleteCommandInput{
		AppID: req.ApplianceId,
		ComID: req.CommandId,
	})

	_res := (<-ch)
	return &res, _res.Err
}

func (h *Handler) NotifyApplianceUpdate(*empty.Empty, aim_api.AimService_NotifyApplianceUpdateServer) error {
	return nil
}
