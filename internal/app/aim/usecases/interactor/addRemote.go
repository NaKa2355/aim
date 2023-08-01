package interactor

import (
	"context"
	"errors"

	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) addRemote(ctx context.Context, _in bdy.AddRemoteInput) (out bdy.AddRemoteOutput, err error) {
	var r *remote.Remote
	switch in := _in.(type) {
	case bdy.AddCustomRemoteInput:
		r, err = remote.NewCustom(in.Name, in.DeviceID)
	case bdy.AddButtonRemoteInput:
		r, err = remote.NewButton(in.Name, in.DeviceID)
	case bdy.AddToggleRemoteInput:
		r, err = remote.NewToggle(in.Name, in.DeviceID)
	case bdy.AddThermostatRemoteInput:
		var scale float64
		switch in.Scale {
		case bdy.Half:
			scale = 0.5
		case bdy.One:
			scale = 1
		}
		r, err = remote.NewThermostat(
			in.Name,
			in.DeviceID,
			scale,
			in.MinimumHeatingTemp,
			in.MaximumHeatingTemp,
			in.MinimumCoolingTemp,
			in.MaximumCoolingTemp,
		)
	default:
		err = bdy.NewError(
			bdy.CodeInvaildInput,
			errors.New("invaild remote type"),
		)
	}

	if err != nil {
		return out, err
	}

	r, err = i.repo.CreateRemote(ctx, r)
	if err != nil {
		return
	}

	out.Remote = bdy.Remote{
		ID:           string(r.ID),
		Name:         string(r.Name),
		Type:         convertType(r.Type),
		DeviceID:     string(r.DeviceID),
		CanAddButton: (r.AddButton() == nil),
	}

	i.output.NotificateRemoteUpdate(
		ctx, bdy.UpdateNotifyOutput{
			Remote: out.Remote,
			Type:   bdy.UpdateTypeAdd,
		},
	)
	return out, err
}

func (i *Interactor) AddRemote(ctx context.Context, in bdy.AddRemoteInput) (bdy.AddRemoteOutput, error) {
	out, err := i.addRemote(ctx, in)
	return out, wrapErr(err)
}
