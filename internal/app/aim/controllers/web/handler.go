package web

import (
	"context"
	"sync"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Boundary interface {
	bdy.ApplianceAdder
	bdy.CommandAdder

	bdy.AppliancesGetter
	bdy.ApplianceGetter
	bdy.IRDataGetter
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
	i            Boundary
	nc           chan aimv1.ApplianceUpdateNotification
	notification aimv1.ApplianceUpdateNotification
	c            *Cond
}

var _ aimv1.AimServiceServer = &Handler{}

func NewHandler() *Handler {
	return &Handler{
		nc: make(chan aimv1.ApplianceUpdateNotification),
		c:  NewCond(&sync.Mutex{}),
	}
}

func (h *Handler) SetInteractor(i Boundary) {
	h.i = i
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
			Scale:              float64(r.Thermostat.Scale),
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
	for i, a := range out.Apps {
		res.Appliances[i] = &aimv1.Appliance{
			Id:            a.ID,
			Name:          a.Name,
			DeviceId:      a.DeviceID,
			CanAddCommand: a.CanAddCommand,
		}
	}
	return
}

func (h *Handler) GetAppliance(ctx context.Context, req *aimv1.GetApplianceRequest) (res *aimv1.GetApplianceResponse, err error) {
	var out bdy.GetApplianceOutput
	res = &aimv1.GetApplianceResponse{}
	in := bdy.GetApplianceInput{
		AppID: req.ApplianceId,
	}

	out, err = h.i.GetAppliance(ctx, in)
	if err != nil {
		return
	}

	res.Appliance = &aimv1.Appliance{
		Id:            out.App.ID,
		Name:          out.App.Name,
		DeviceId:      out.App.DeviceID,
		CanAddCommand: out.App.CanAddCommand,
	}
	return res, err
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
			Id:        c.ID,
			Name:      c.Name,
			CanRename: c.CanRename,
			CanDelete: c.CanDelete,
		}
	}

	return
}

func (h *Handler) GetIrData(ctx context.Context, req *aimv1.GetIrDataRequest) (res *anypb.Any, err error) {
	var in bdy.GetIRDataInput
	var out bdy.GetIRDataOutput
	res = &anypb.Any{}

	in.AppID = req.ApplianceId
	in.ComID = req.CommandId
	out, err = h.i.GetIRData(ctx, in)
	if err != nil {
		return
	}

	proto.Unmarshal(out.IRData, res)
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

func (h *Handler) NotificateApplianceUpdate(ctx context.Context, o bdy.UpdateNotifyOutput) {
	defer h.c.L.Unlock()
	var updateType aimv1.ApplianceUpdateNotification_UpdateType
	switch o.Type {
	case bdy.UpdateTypeAdd:
		updateType = aimv1.ApplianceUpdateNotification_UPDATE_TYPE_ADD
	case bdy.UpdateTypeDelete:
		updateType = aimv1.ApplianceUpdateNotification_UPDATE_TYPE_DELETE
	}

	h.c.L.Lock()
	h.notification = aimv1.ApplianceUpdateNotification{
		ApplianceId: o.AppID,
		UpdateType:  updateType,
	}
	h.c.Broadcast()
}

func (h *Handler) NotifyApplianceUpdate(_ *empty.Empty, stream aimv1.AimService_NotifyApplianceUpdateServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case <-h.c.NotifyChan():
			h.c.L.Lock()
			err := stream.Send(&h.notification)
			h.c.L.Unlock()
			if err != nil {
				return err
			}
		}
	}
}
