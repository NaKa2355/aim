package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) addRemote(ctx context.Context, in bdy.AddRemoteInput) (out bdy.AddRemoteOutput, err error) {
	buttons := make([]*button.Button, len(in.Buttons))

	for i, b := range in.Buttons {
		buttons[i] = button.New(button.Name(b.Name), button.Tag(b.Tag), irdata.IRData{})
	}
	r, err := remote.NewRemote(in.Name, in.DeviceID, in.Tag, buttons)

	if err != nil {
		return out, err
	}

	r, err = i.repo.CreateRemote(ctx, r)
	if err != nil {
		return
	}

	out.Remote = bdy.Remote{
		ID:       string(r.ID),
		Name:     string(r.Name),
		Tag:      string(r.Tag),
		DeviceID: string(r.DeviceID),
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
