package web

import (
	"context"
	"errors"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	v1 "github.com/NaKa2355/irdeck-proto/gen/go/common/irdata/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/proto"
)

type Boundary interface {
	bdy.ApplianceAdder
	bdy.CommandAdder

	bdy.AppliancesGetter
	bdy.CommandGetter
	bdy.CommandsGetter

	bdy.ApplianceRenamer
	bdy.IRDeviceChanger
	bdy.CommandRenamer
	bdy.IRDataSetter

	bdy.ApplianceDeleter
	bdy.CommandDeleter
}

type Handler struct {
	aimv1.UnimplementedAimServiceServer
	i Boundary
}

var _ aimv1.AimServiceServer = &Handler{}

func NewHandler(i Boundary) *Handler {
	return &Handler{
		i: i,
	}
}

func (h *Handler) AddAppliance(ctx context.Context, _req *aimv1.AddApplianceRequest) (res *aimv1.AddAppResponse, err error) {
	var in bdy.AddApplianceInput
	var out bdy.AddAppOutput

	switch r := _req.GetAppliance().(type) {
	case *aimv1.AddApplianceRequest_Custom:
		in = bdy.AddCustomInput{
			Name:     r.Custom.Name,
			DeviceID: r.Custom.DeviceId,
		}
	case *aimv1.AddApplianceRequest_Button:
		in = bdy.AddButtonInput{
			Name:     r.Button.Name,
			DeviceID: r.Button.DeviceId,
		}
	case *aimv1.AddApplianceRequest_Toggle:
		in = bdy.AddToggleInput{
			Name:     r.Toggle.Name,
			DeviceID: r.Toggle.DeviceId,
		}
	case *aimv1.AddApplianceRequest_Thermostat:
		in = bdy.AddThermostatInput{
			Name:               r.Thermostat.Name,
			DeviceID:           r.Thermostat.DeviceId,
			Scale:              float64(r.Thermostat.TempScale),
			MinimumHeatingTemp: int(r.Thermostat.MinimumHeatingTemp),
			MaximumHeatingTemp: int(r.Thermostat.MaximumHeatingTemp),
			MinimumCoolingTemp: int(r.Thermostat.MinimumCoolingTemp),
			MaximumCoolingTemp: int(r.Thermostat.MaximumCoolingTemp),
		}
	}

	out, err = h.i.AddAppliance(ctx, in)
	if err != nil {
		return
	}

	res = &aimv1.AddAppResponse{
		ApplianceId: out.ID,
	}
	return
}

func (h *Handler) AddCommand(ctx context.Context, req *aimv1.AddCommandRequest) (e *empty.Empty, err error) {
	e = &empty.Empty{}

	in := bdy.AddCommandInput{
		AppID: req.ApplianceId,
		Name:  req.Name,
	}
	err = h.i.AddCommand(ctx, in)
	return
}

func (h *Handler) GetAppliances(ctx context.Context, _ *empty.Empty) (res *aimv1.GetAppliancesResponse, err error) {
	var out bdy.GetAppliancesOutput
	res = &aimv1.GetAppliancesResponse{}

	out, err = h.i.GetAppliances(ctx)
	if err != nil {
		return
	}

	res.Appliances = make([]*aimv1.Appliance, len(out.Apps))
	for i, _a := range out.Apps {
		res.Appliances[i] = &aimv1.Appliance{}
		app := res.Appliances[i]

		switch a := _a.(type) {
		case bdy.Custom:
			app.Appliance = &aimv1.Appliance_Custom{
				Custom: &aimv1.Custom{
					Id:       a.ID,
					Name:     a.Name,
					DeviceId: a.DeviceID,
				},
			}

		case bdy.Button:
			app.Appliance = &aimv1.Appliance_Button{
				Button: &aimv1.Button{
					Id:       a.ID,
					Name:     a.Name,
					DeviceId: a.DeviceID,
				},
			}

		case bdy.Toggle:
			app.Appliance = &aimv1.Appliance_Toggle{
				Toggle: &aimv1.Toggle{
					Id:       a.ID,
					Name:     a.Name,
					DeviceId: a.DeviceID,
				},
			}

		case bdy.Thermostat:
			app.Appliance = &aimv1.Appliance_Thermostat{
				Thermostat: &aimv1.Thermostat{
					Id:                 a.ID,
					Name:               a.Name,
					DeviceId:           a.DeviceID,
					TempScale:          float32(a.Scale),
					MinimumHeatingTemp: uint32(a.MinimumHeatingTemp),
					MaximumHeatingTemp: uint32(a.MaximumHeatingTemp),
					MinimumCoolingTemp: uint32(a.MinimumCoolingTemp),
					MaximumCoolingTemp: uint32(a.MaximumCoolingTemp),
				},
			}

		default:
			err = errors.New("undefined appliance type")
		}
	}
	return
}

func (h *Handler) GetCommands(ctx context.Context, req *aimv1.GetCommandsRequest) (res *aimv1.GetCommandsResponse, err error) {
	var in = bdy.GetCommandsInput{}
	var out bdy.GetCommandsOutput
	res = &aimv1.GetCommandsResponse{}

	in.AppID = req.ApplianceId

	out, err = h.i.GetCommands(ctx, in)
	if err != nil {
		return
	}

	res.Commands = make([]*aimv1.Command, len(out.Commands))
	for i, c := range out.Commands {
		res.Commands[i] = &aimv1.Command{
			Id:   c.ID,
			Name: c.Name,
		}
	}

	return
}

func (h *Handler) GetIrData(ctx context.Context, req *aimv1.GetIrDataRequest) (res *v1.IrData, err error) {
	var in bdy.GetCommandInput
	var out bdy.GetCommandOutput
	res = &v1.IrData{}

	in.AppID = req.ApplianceId
	in.ComID = req.CommandId
	out, err = h.i.GetCommand(ctx, in)
	if err != nil {
		return
	}

	proto.Unmarshal(out.Data, res)
	return
}

func (h *Handler) RenameAppliance(ctx context.Context, req *aimv1.RenameApplianceRequest) (e *empty.Empty, err error) {
	var in bdy.RenameAppInput
	e = &empty.Empty{}

	in.AppID = req.ApplianceId
	in.Name = req.Name
	err = h.i.RenameAppliance(ctx, in)
	return
}

func (h *Handler) ChangeDevice(ctx context.Context, req *aimv1.ChangeDeviceRequest) (e *empty.Empty, err error) {
	var in bdy.ChangeIRDevInput
	e = &empty.Empty{}

	in.AppID = req.ApplianceId
	in.DeviceID = req.DeviceId
	err = h.i.ChangeIRDevice(ctx, in)
	return
}

func (h *Handler) RenameCommand(ctx context.Context, req *aimv1.RenameCommandRequest) (e *empty.Empty, err error) {
	var in bdy.RenameCommandInput
	e = &empty.Empty{}

	in.AppID = req.ApplianceId
	in.ComID = req.CommandId
	in.Name = req.Name
	err = h.i.RenameCommand(ctx, in)
	return
}

func (h *Handler) SetIrData(ctx context.Context, req *aimv1.SetIRDataRequest) (e *empty.Empty, err error) {
	var in bdy.SetIRDataInput
	e = &empty.Empty{}

	in.AppID = req.ApplianceId
	in.ComID = req.CommandId
	in.Data, err = proto.Marshal(req.Irdata)
	if err != nil {
		return
	}

	err = h.i.SetIRData(ctx, in)
	return
}

func (h *Handler) DeleteAppliance(ctx context.Context, req *aimv1.DeleteApplianceRequest) (e *empty.Empty, err error) {
	var in bdy.DeleteAppInput
	e = &empty.Empty{}

	in.AppID = req.ApplianceId

	err = h.i.DeleteAppliance(ctx, in)
	return
}

func (h *Handler) DeleteCommand(ctx context.Context, req *aimv1.DeleteCommandRequest) (e *empty.Empty, err error) {
	var in bdy.DeleteCommandInput
	e = &empty.Empty{}

	in.AppID = req.ApplianceId
	in.ComID = req.CommandId

	err = h.i.DeleteCommand(ctx, in)
	return
}

func (h *Handler) NotifyApplianceUpdate(_ *empty.Empty, req aimv1.AimService_NotifyApplianceUpdateServer) error {
	return nil
}

func convertCommands(coms []bdy.Command) []*aimv1.Command {
	res := make([]*aimv1.Command, len(coms))
	for i, c := range coms {
		res[i] = &aimv1.Command{
			Id:   c.ID,
			Name: c.Name,
		}
	}
	return res
}
