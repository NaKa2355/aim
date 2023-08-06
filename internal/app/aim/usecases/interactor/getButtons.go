package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) getButtons(ctx context.Context, in bdy.GetButtonsInput) (out bdy.GetButtonsOutput, err error) {
	var buttons []*button.Button
	_, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return
	}

	buttons, err = i.repo.ReadButtons(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return
	}

	out.Buttons = make([]bdy.Button, len(buttons))
	for i, b := range buttons {
		out.Buttons[i].ID = string(b.ID)
		out.Buttons[i].Name = string(b.Name)
		out.Buttons[i].Tag = string(b.Tag)
		out.Buttons[i].HasIRData = (len(b.IRData) != 0)
	}

	return
}

func (i *Interactor) GetButtons(ctx context.Context, in bdy.GetButtonsInput) (bdy.GetButtonsOutput, error) {
	out, err := i.getButtons(ctx, in)
	return out, wrapErr(err)
}
