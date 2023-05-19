package interactor

import (
	"context"
	"errors"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	remote "github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func convertType(in remote.RemoteType) (out bdy.RemoteType) {
	switch in {
	case remote.TypeCustom:
		return bdy.TypeCustom
	case remote.TypeButton:
		return bdy.TypeButton
	case remote.TypeToggle:
		return bdy.TypeToggle
	case remote.TypeThermostat:
		return bdy.TypeThermostat
	default:
		return bdy.RemoteType(in)
	}
}

func (i *Interactor) addRemote(ctx context.Context, _in bdy.AddRemoteInput) (out bdy.AddRemoteOutput, err error) {
	var r *remote.Remote
	switch in := _in.(type) {
	case bdy.AddCustomRemoteInput:
		r, err = remote.NewCustom(in.Name, in.DeviceID)
	case bdy.AddButtonRemoteInput:
		r, err = remote.NewButton(in.Name, in.DeviceID)
	case bdy.AddToggleRemoteInput:
		r, err = remote.NewToggle(in.Name, in.DeviceID)
	case bdy.AddThermostatRemoteInput:
		r, err = remote.NewThermostat(
			in.Name,
			in.DeviceID,
			in.Scale,
			in.MinimumHeatingTemp,
			in.MaximumHeatingTemp,
			in.MinimumCoolingTemp,
			in.MaximumCoolingTemp,
		)
	default:
		err = errors.New("invaild input")
	}
	if err != nil {
		return out, err
	}

	r, err = i.repo.CreateRemote(ctx, r)
	if err != nil {
		return
	}

	i.output.NotificateRemoteUpdate(
		ctx, bdy.UpdateNotifyOutput{
			RemoteID: string(r.ID),
			Type:     bdy.UpdateTypeAdd,
		},
	)
	out.Remote = bdy.Remote{
		ID:           string(r.ID),
		Name:         string(r.Name),
		Type:         convertType(r.Type),
		DeviceID:     string(r.DeviceID),
		CanAddButton: (r.AddButton() == nil),
	}
	return out, err
}

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

func (i *Interactor) getRemotes(ctx context.Context) (out bdy.GetRemotesOutput, err error) {
	var remotes []*remote.Remote

	remotes, err = i.repo.ReadRemotes(ctx)
	if err != nil {
		return
	}

	out.Remotes = make([]bdy.Remote, len(remotes))
	for i, r := range remotes {
		out.Remotes[i] = bdy.Remote{
			ID:           string(r.ID),
			DeviceID:     string(r.DeviceID),
			Type:         convertType(r.Type),
			Name:         string(r.Name),
			CanAddButton: (r.AddButton() == nil),
		}
	}
	return
}

func (i *Interactor) getRemote(ctx context.Context, in bdy.GetRemoteInput) (out bdy.GetRemoteOutput, err error) {
	var r *remote.Remote

	r, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return out, err
	}

	out.Remote = bdy.Remote{
		ID:           string(r.ID),
		DeviceID:     string(r.DeviceID),
		Type:         convertType(r.Type),
		Name:         string(r.Name),
		CanAddButton: (r.AddButton() == nil),
	}
	return out, err
}

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

func (i *Interactor) getIRData(ctx context.Context, in bdy.GetIRDataInput) (out bdy.GetIRDataOutput, err error) {
	var b *button.Button

	b, err = i.repo.ReadButton(ctx, remote.ID(in.RemoteID), button.ID(in.ButtonID))
	if err != nil {
		return
	}

	out.IRData = bdy.IRData(b.IRData)

	return
}

func (i *Interactor) editRemote(ctx context.Context, in bdy.EditRemoteInput) (err error) {
	var r *remote.Remote

	r, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return
	}

	err = r.SetName(in.Name)
	if err != nil {
		return
	}

	err = r.SetDeviceID(in.DeviceID)
	if err != nil {
		return
	}

	err = i.repo.UpdateRemote(ctx, r)
	return
}

func (i *Interactor) renameButton(ctx context.Context, in bdy.EditButtonInput) (err error) {
	var r *remote.Remote
	var b *button.Button

	r, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return
	}

	err = r.ChangeButtonName()
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

// Delete
func (i *Interactor) deleteRemote(ctx context.Context, in bdy.DeleteRemoteInput) (err error) {
	if _, err = i.repo.ReadRemote(ctx, remote.ID(in.RemoteID)); err != nil {
		return err
	}

	err = i.repo.DeleteRemote(ctx, remote.ID(in.RemoteID))
	if err != nil {
		return err
	}

	i.output.NotificateRemoteUpdate(
		ctx,
		bdy.UpdateNotifyOutput{
			RemoteID: in.RemoteID,
			Type:     bdy.UpdateTypeDelete,
		},
	)
	return err
}

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
