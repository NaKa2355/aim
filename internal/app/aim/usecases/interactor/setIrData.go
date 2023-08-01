package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) setIRData(ctx context.Context, in bdy.SetIRDataInput) (err error) {
	var b *button.Button

	b, err = i.repo.ReadButton(ctx, remote.ID(in.RemoteID), button.ID(in.ButtonID))
	if err != nil {
		return
	}

	b.SetRawIRData(irdata.IRData(in.Data))
	err = i.repo.UpdateButton(ctx, remote.ID(in.RemoteID), b)
	return
}

func (i *Interactor) SetIRData(ctx context.Context, in bdy.SetIRDataInput) error {
	err := i.setIRData(ctx, in)
	return wrapErr(err)
}
