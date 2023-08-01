package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
)

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
			Scale:              convertThermostatScale(r.Thermostat.Scale),
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
