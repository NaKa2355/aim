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
	bdy.RemoteAdder
	bdy.ButtonAdder

	bdy.RemotesGetter
	bdy.RemoteGetter
	bdy.IRDataGetter
	bdy.ButtonsGetter

	bdy.RemoteEditor
	bdy.ButtonEditor
	bdy.IRDataSetter

	bdy.RemoteDeleter
	bdy.ButtonDeleter
}

type Handler struct {
	aimv1.UnimplementedAimServiceServer
	i            Boundary
	nc           chan aimv1.RemoteUpdateNotification
	notification aimv1.RemoteUpdateNotification
	c            *Cond
}

var _ aimv1.AimServiceServer = &Handler{}

func NewHandler() *Handler {
	return &Handler{
		nc: make(chan aimv1.RemoteUpdateNotification),
		c:  NewCond(&sync.Mutex{}),
	}
}

func convertRemoteType(in bdy.RemoteType) (out aimv1.Remote_RemoteType) {
	switch in {
	case bdy.TypeCustom:
		return aimv1.Remote_CUSTOM
	case bdy.TypeButton:
		return aimv1.Remote_BUTTON
	case bdy.TypeToggle:
		return aimv1.Remote_TOGGLE
	case bdy.TypeThermostat:
		return aimv1.Remote_THERMOSTAT
	default:
		return aimv1.Remote_REMOTE_TYPE_UNKNOWN
	}
}

func (h *Handler) SetInteractor(i Boundary) {
	h.i = i
}

func (h *Handler) AddRemote(ctx context.Context, _req *aimv1.AddRemoteRequest) (res *aimv1.AddRemoteResponse, err error) {
	var in bdy.AddRemoteInput
	var out bdy.AddRemoteOutput

	switch r := _req.GetRemote().(type) {
	case *aimv1.AddRemoteRequest_Custom:
		in = bdy.AddCustomRemoteInput{
			Name:     r.Custom.Name,
			DeviceID: r.Custom.DeviceId,
		}
	case *aimv1.AddRemoteRequest_Button:
		in = bdy.AddButtonRemoteInput{
			Name:     r.Button.Name,
			DeviceID: r.Button.DeviceId,
		}
	case *aimv1.AddRemoteRequest_Toggle:
		in = bdy.AddToggleRemoteInput{
			Name:     r.Toggle.Name,
			DeviceID: r.Toggle.DeviceId,
		}
	case *aimv1.AddRemoteRequest_Thermostat:
		in = bdy.AddThermostatRemoteInput{
			Name:               r.Thermostat.Name,
			DeviceID:           r.Thermostat.DeviceId,
			Scale:              float64(r.Thermostat.Scale),
			MinimumHeatingTemp: int(r.Thermostat.MinimumHeatingTemp),
			MaximumHeatingTemp: int(r.Thermostat.MaximumHeatingTemp),
			MinimumCoolingTemp: int(r.Thermostat.MinimumCoolingTemp),
			MaximumCoolingTemp: int(r.Thermostat.MaximumCoolingTemp),
		}
	}

	out, err = h.i.AddRemote(ctx, in)
	if err != nil {
		return
	}

	res = &aimv1.AddRemoteResponse{
		Remote: &aimv1.Remote{
			Id:           out.Remote.ID,
			Name:         out.Remote.Name,
			DeviceId:     out.Remote.DeviceID,
			RemoteType:   convertRemoteType(out.Remote.Type),
			CanAddButton: out.Remote.CanAddButton,
		},
	}
	return
}

func (h *Handler) AddButton(ctx context.Context, req *aimv1.AddButtonRequest) (e *empty.Empty, err error) {
	e = &empty.Empty{}

	in := bdy.AddButtonInput{
		RemoteID: req.RemoteId,
		Name:     req.Name,
	}
	err = h.i.AddButton(ctx, in)
	return
}

func (h *Handler) GetRemotes(ctx context.Context, _ *empty.Empty) (res *aimv1.GetRemotesResponse, err error) {
	var out bdy.GetRemotesOutput
	res = &aimv1.GetRemotesResponse{}

	out, err = h.i.GetRemotes(ctx)
	if err != nil {
		return
	}

	res.Remotes = make([]*aimv1.Remote, len(out.Remotes))
	for i, r := range out.Remotes {
		res.Remotes[i] = &aimv1.Remote{
			Id:           r.ID,
			Name:         r.Name,
			RemoteType:   convertRemoteType(r.Type),
			DeviceId:     r.DeviceID,
			CanAddButton: r.CanAddButton,
		}
	}
	return
}

func (h *Handler) GetRemote(ctx context.Context, req *aimv1.GetRemoteRequest) (res *aimv1.GetRemoteResponse, err error) {
	var out bdy.GetRemoteOutput
	res = &aimv1.GetRemoteResponse{}
	in := bdy.GetRemoteInput{
		RemoteID: req.RemoteId,
	}

	out, err = h.i.GetRemote(ctx, in)
	if err != nil {
		return
	}

	res.Remote = &aimv1.Remote{
		Id:           out.Remote.ID,
		Name:         out.Remote.Name,
		RemoteType:   convertRemoteType(out.Remote.Type),
		DeviceId:     out.Remote.DeviceID,
		CanAddButton: out.Remote.CanAddButton,
	}
	return res, err
}

func (h *Handler) GetButtons(ctx context.Context, req *aimv1.GetButtonsRequest) (res *aimv1.GetButtonsResponse, err error) {
	var in = bdy.GetButtonsInput{}
	var out bdy.GetButtonsOutput
	res = &aimv1.GetButtonsResponse{}

	in.RemoteID = req.RemoteId

	out, err = h.i.GetButtons(ctx, in)
	if err != nil {
		return
	}

	res.Buttons = make([]*aimv1.Button, len(out.Buttons))
	for i, b := range out.Buttons {
		res.Buttons[i] = &aimv1.Button{
			Id:        b.ID,
			Name:      b.Name,
			CanRename: b.CanRename,
			CanDelete: b.CanDelete,
			HasIrdata: b.HasIRData,
		}
	}

	return
}

func (h *Handler) GetIrData(ctx context.Context, req *aimv1.GetIrDataRequest) (res *anypb.Any, err error) {
	var in bdy.GetIRDataInput
	var out bdy.GetIRDataOutput
	res = &anypb.Any{}

	in.RemoteID = req.RemoteId
	in.ButtonID = req.ButtonId
	out, err = h.i.GetIRData(ctx, in)
	if err != nil {
		return
	}

	proto.Unmarshal(out.IRData, res)
	return
}

func (h *Handler) EditRemote(ctx context.Context, req *aimv1.EditRemoteRequest) (e *empty.Empty, err error) {
	var in bdy.EditRemoteInput
	e = &empty.Empty{}

	in.RemoteID = req.RemoteId
	in.Name = req.Name
	in.DeviceID = req.DeviceId
	err = h.i.EditRemote(ctx, in)
	return
}

func (h *Handler) EditButton(ctx context.Context, req *aimv1.EditButtonRequest) (e *empty.Empty, err error) {
	var in bdy.EditButtonInput
	e = &empty.Empty{}

	in.RemoteID = req.RemoteId
	in.ButtonID = req.ButtonId
	in.Name = req.Name
	err = h.i.EditButton(ctx, in)
	return
}

func (h *Handler) SetIrData(ctx context.Context, req *aimv1.SetIRDataRequest) (e *empty.Empty, err error) {
	var in bdy.SetIRDataInput
	e = &empty.Empty{}

	in.RemoteID = req.RemoteId
	in.ButtonID = req.ButtonId
	in.Data, err = proto.Marshal(req.Irdata)
	if err != nil {
		return
	}

	err = h.i.SetIRData(ctx, in)
	return
}

func (h *Handler) DeleteRemote(ctx context.Context, req *aimv1.DeleteRemoteRequest) (e *empty.Empty, err error) {
	var in bdy.DeleteRemoteInput
	e = &empty.Empty{}

	in.RemoteID = req.RemoteId

	err = h.i.DeleteRemote(ctx, in)
	return
}

func (h *Handler) DeleteButton(ctx context.Context, req *aimv1.DeleteButtonRequest) (e *empty.Empty, err error) {
	var in bdy.DeleteButtonInput
	e = &empty.Empty{}

	in.RemoteID = req.RemoteId
	in.ButtonID = req.ButtonId
	err = h.i.DeleteButton(ctx, in)
	return
}

func (h *Handler) NotificateRemoteUpdate(ctx context.Context, o bdy.UpdateNotifyOutput) {
	defer h.c.L.Unlock()
	var updateType aimv1.RemoteUpdateNotification_UpdateType
	switch o.Type {
	case bdy.UpdateTypeAdd:
		updateType = aimv1.RemoteUpdateNotification_UPDATE_TYPE_ADD
	case bdy.UpdateTypeDelete:
		updateType = aimv1.RemoteUpdateNotification_UPDATE_TYPE_DELETE
	}

	h.c.L.Lock()
	h.notification = aimv1.RemoteUpdateNotification{
		RemoteId:   o.RemoteID,
		UpdateType: updateType,
	}
	h.c.Broadcast()
}

func (h *Handler) NotifyRemoteUpdate(_ *empty.Empty, stream aimv1.AimService_NotifyRemoteUpdateServer) error {
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
