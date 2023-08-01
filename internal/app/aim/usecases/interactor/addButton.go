package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) addButton(ctx context.Context, in bdy.AddButtonInput) (err error) {
	var r *remote.Remote
	var b *button.Button

	r, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return
	}

	err = r.AddButton()
	if err != nil {
		return
	}

	b = button.New(button.Name(in.Name), irdata.IRData{})
	_, err = i.repo.CreateButton(ctx, remote.ID(in.RemoteID), b)
	return
}

func (i *Interactor) AddButton(ctx context.Context, in bdy.AddButtonInput) error {
	err := i.addButton(ctx, in)
	return wrapErr(err)
}
