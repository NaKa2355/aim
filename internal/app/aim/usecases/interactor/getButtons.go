package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) getButtons(ctx context.Context, in bdy.GetButtonsInput) (out bdy.GetButtonsOutput, err error) {
	var buttons []*button.Button
	r, err := i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return
	}

	buttons, err = i.repo.ReadButtons(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return
	}

	out.Buttons = make([]bdy.Button, len(buttons))
	for i, c := range buttons {
		c.GetRawIRData()
		out.Buttons[i].ID = string(c.ID)
		out.Buttons[i].Name = string(c.Name)
		out.Buttons[i].CanRename = (r.ChangeButtonName() == nil)
		out.Buttons[i].CanDelete = (r.ChangeButtonName() == nil)
		out.Buttons[i].HasIRData = (len(c.IRData) != 0)
	}

	return
}

func (i *Interactor) GetButtons(ctx context.Context, in bdy.GetButtonsInput) (bdy.GetButtonsOutput, error) {
	out, err := i.getButtons(ctx, in)
	return out, wrapErr(err)
}
