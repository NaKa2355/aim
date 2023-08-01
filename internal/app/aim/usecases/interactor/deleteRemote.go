package interactor

import (
	"context"

	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

// Delete
func (i *Interactor) deleteRemote(ctx context.Context, in bdy.DeleteRemoteInput) (err error) {
	r, err := i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return err
	}

	err = i.repo.DeleteRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return err
	}

	i.output.NotificateRemoteUpdate(
		ctx,
		bdy.UpdateNotifyOutput{
			Remote: bdy.Remote{
				ID:           string(r.ID),
				Name:         string(r.Name),
				Type:         bdy.RemoteType(r.Type),
				DeviceID:     string(r.DeviceID),
				CanAddButton: r.Type == remote.TypeCustom,
			},
			Type: bdy.UpdateTypeDelete,
		},
	)
	return err
}

// Delete
func (i *Interactor) DeleteRemote(ctx context.Context, in bdy.DeleteRemoteInput) error {
	err := i.deleteRemote(ctx, in)
	return wrapErr(err)
}
