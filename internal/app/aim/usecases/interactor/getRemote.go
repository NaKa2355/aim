package interactor

import (
	"context"

	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) getRemote(ctx context.Context, in bdy.GetRemoteInput) (out bdy.GetRemoteOutput, err error) {
	var r *remote.Remote

	r, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return out, err
	}

	out.Remote = bdy.Remote{
		ID:           string(r.ID),
		DeviceID:     string(r.DeviceID),
		Type:         convertType(r.Type),
		Name:         string(r.Name),
		CanAddButton: (r.AddButton() == nil),
	}
	return out, err
}

func (i *Interactor) GetRemote(ctx context.Context, in bdy.GetRemoteInput) (bdy.GetRemoteOutput, error) {
	out, err := i.getRemote(ctx, in)
	return out, wrapErr(err)
}
