package interactor

import (
	"context"

	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) editRemote(ctx context.Context, in bdy.EditRemoteInput) (err error) {
	var r *remote.Remote

	r, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return
	}

	err = r.SetName(in.Name)
	if err != nil {
		return
	}

	err = r.SetDeviceID(in.DeviceID)
	if err != nil {
		return
	}

	err = i.repo.UpdateRemote(ctx, r)
	i.output.NotificateRemoteUpdate(ctx, bdy.UpdateNotifyOutput{
		Type: bdy.UpdateTypeUpdate,
		Remote: bdy.Remote{
			ID:       string(r.ID),
			Name:     string(r.Name),
			Tag:      string(r.Tag),
			DeviceID: string(r.DeviceID),
		},
	})
	return
}

// Update
func (i *Interactor) EditRemote(ctx context.Context, in bdy.EditRemoteInput) error {
	err := i.editRemote(ctx, in)
	return wrapErr(err)
}
