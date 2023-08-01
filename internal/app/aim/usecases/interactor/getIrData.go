package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) getIRData(ctx context.Context, in bdy.GetIRDataInput) (out bdy.GetIRDataOutput, err error) {
	var b *button.Button

	b, err = i.repo.ReadButton(ctx, remote.ID(in.RemoteID), button.ID(in.ButtonID))
	if err != nil {
		return
	}

	out.IRData = bdy.IRData(b.IRData)

	return
}

func (i *Interactor) GetIRData(ctx context.Context, in bdy.GetIRDataInput) (bdy.GetIRDataOutput, error) {
	out, err := i.getIRData(ctx, in)
	return out, wrapErr(err)
}
