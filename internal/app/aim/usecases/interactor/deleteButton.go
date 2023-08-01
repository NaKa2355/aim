package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) deleteButton(ctx context.Context, in bdy.DeleteButtonInput) (err error) {
	var r *remote.Remote

	r, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return
	}

	err = r.RemoveButton()
	if err != nil {
		return
	}

	err = i.repo.DeleteButton(ctx, remote.ID(in.RemoteID), button.ID(in.ButtonID))
	return
}

func (i *Interactor) DeleteButton(ctx context.Context, in bdy.DeleteButtonInput) error {
	err := i.deleteButton(ctx, in)
	return wrapErr(err)
}
