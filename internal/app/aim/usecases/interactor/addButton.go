package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) addButton(ctx context.Context, in bdy.AddButtonInput) (out bdy.AddButtonOutput, err error) {
	_, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return out, err
	}
	b := button.New(button.Name(in.Name), button.Tag(in.Tag), irdata.IRData{})

	b, err = i.repo.CreateButton(ctx, remote.ID(in.RemoteID), b)

	out.Button = bdy.Button{
		ID:        string(b.ID),
		Name:      string(b.Name),
		Tag:       string(b.Tag),
		HasIRData: (len(b.IRData) == 0),
	}
	return
}

func (i *Interactor) AddButton(ctx context.Context, in bdy.AddButtonInput) (out bdy.AddButtonOutput, err error) {
	out, err = i.addButton(ctx, in)
	err = wrapErr(err)
	return out, err
}
