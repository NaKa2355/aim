package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) renameButton(ctx context.Context, in bdy.EditButtonInput) (err error) {
	var b *button.Button

	_, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return
	}

	b, err = i.repo.ReadButton(ctx, remote.ID(in.RemoteID), button.ID(in.ButtonID))
	if err != nil {
		return
	}

	b.SetName(button.Name(in.Name))
	err = i.repo.UpdateButton(ctx, remote.ID(in.RemoteID), b)

	return
}

func (i *Interactor) EditButton(ctx context.Context, in bdy.EditButtonInput) error {
	err := i.renameButton(ctx, in)
	return wrapErr(err)
}
