package interactor

import (
	"context"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) renameButton(ctx context.Context, in bdy.EditButtonInput) (err error) {
	var b *button.Button

	b, err = i.repo.ReadButton(ctx, remote.ID(in.RemoteID), button.ID(in.ButtonID))
	if err != nil {
		return
	}

	if b.Tag != "" {
		return bdy.NewError(bdy.CodeInvaildOperation, fmt.Errorf("cannot edit a button which has tag"))
	}

	b.SetName(button.Name(in.Name))
	err = i.repo.UpdateButton(ctx, remote.ID(in.RemoteID), b)

	return
}

func (i *Interactor) EditButton(ctx context.Context, in bdy.EditButtonInput) error {
	err := i.renameButton(ctx, in)
	return wrapErr(err)
}
