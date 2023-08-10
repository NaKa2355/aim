package interactor

import (
	"context"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) deleteButton(ctx context.Context, in bdy.DeleteButtonInput) (err error) {
	b, err := i.repo.ReadButton(ctx, remote.ID(in.RemoteID), button.ID(in.ButtonID))
	if err != nil {
		return
	}

	if b.Tag != "" {
		return bdy.NewError(bdy.CodeInvaildOperation, fmt.Errorf("cannot delete a button which has tag"))
	}

	err = i.repo.DeleteButton(ctx, remote.ID(in.RemoteID), button.ID(in.ButtonID))
	return
}

func (i *Interactor) DeleteButton(ctx context.Context, in bdy.DeleteButtonInput) error {
	err := i.deleteButton(ctx, in)
	return wrapErr(err)
}
