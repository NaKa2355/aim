package interactor

import (
	"context"
	"errors"

	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) addAppliance(ctx context.Context, _in bdy.AddApplianceInput) (out bdy.AddAppOutput, err error) {
	var a *app.Appliance
	switch in := _in.(type) {
	case bdy.AddCustomInput:
		a, err = app.NewCustom(in.Name, in.DeviceID)
	case bdy.AddButtonInput:
		a, err = app.NewButton(in.Name, in.DeviceID)
	case bdy.AddToggleInput:
		a, err = app.NewToggle(in.Name, in.DeviceID)
	case bdy.AddThermostatInput:
		a, err = app.NewThermostat(
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

	a, err = i.repo.CreateAppliance(ctx, a)
	if err != nil {
		return
	}

	i.output.NotificateApplianceUpdate(
		ctx, bdy.UpdateNotifyOutput{
			AppID: string(a.ID),
			Type:  bdy.UpdateTypeAdd,
		},
	)
	out.ID = string(a.ID)
	return out, err
}

func (i *Interactor) addCommand(ctx context.Context, in bdy.AddCommandInput) (err error) {
	var a *app.Appliance
	var com command.Command

	a, err = i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	err = a.AddCommand()
	if err != nil {
		return
	}

	com = command.New(command.Name(in.Name), irdata.IRData{})
	_, err = i.repo.CreateCommand(ctx, app.ID(in.AppID), com)
	return
}

func (i *Interactor) getAppliances(ctx context.Context) (out bdy.GetAppliancesOutput, err error) {
	var apps []*app.Appliance

	apps, err = i.repo.ReadApps(ctx)
	if err != nil {
		return
	}

	out.Apps = make([]bdy.Appliance, len(apps))
	for i, a := range apps {
		out.Apps[i] = bdy.Appliance{
			ID:            string(a.ID),
			DeviceID:      string(a.DeviceID),
			Name:          string(a.Name),
			CanAddCommand: (a.AddCommand() == nil),
		}
	}
	return
}

func (i *Interactor) getAppliance(ctx context.Context, in bdy.GetApplianceInput) (out bdy.GetApplianceOutput, err error) {
	var a *app.Appliance

	a, err = i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		return out, err
	}

	out.App = bdy.Appliance{
		ID:            string(a.ID),
		DeviceID:      string(a.DeviceID),
		Name:          string(a.Name),
		CanAddCommand: (a.AddCommand() == nil),
	}
	return out, err
}

func (i *Interactor) getCommands(ctx context.Context, in bdy.GetCommandsInput) (out bdy.GetCommandsOutput, err error) {
	var coms []command.Command
	a, err := i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	coms, err = i.repo.ReadCommands(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	out.Commands = make([]bdy.Command, len(coms))
	for i, c := range coms {
		out.Commands[i].ID = string(c.ID)
		out.Commands[i].Name = string(c.Name)
		out.Commands[i].CanRename = (a.ChangeCommandName() == nil)
		out.Commands[i].CanDelete = (a.RemoveCommand() == nil)
	}

	return
}

func (i *Interactor) getIRData(ctx context.Context, in bdy.GetIRDataInput) (out bdy.GetIRDataOutput, err error) {
	var com command.Command

	com, err = i.repo.ReadCommand(ctx, app.ID(in.AppID), command.ID(in.ComID))
	if err != nil {
		return
	}

	out.IRData = bdy.IRData(com.IRData)

	return
}

// Update
func (i *Interactor) renameAppliance(ctx context.Context, in bdy.RenameAppInput) (err error) {
	var a *app.Appliance

	a, err = i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	err = a.SetName(in.Name)
	if err != nil {
		return
	}

	err = i.repo.UpdateApp(ctx, a)
	return
}

func (i *Interactor) changeIRDevice(ctx context.Context, in bdy.ChangeIRDevInput) (err error) {
	var a *app.Appliance

	a, err = i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	a.SetDeviceID(in.DeviceID)
	err = i.repo.UpdateApp(ctx, a)

	return
}

func (i *Interactor) renameCommand(ctx context.Context, in bdy.RenameCommandInput) (err error) {
	var a *app.Appliance
	var c command.Command

	a, err = i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	err = a.ChangeCommandName()
	if err != nil {
		return
	}

	c, err = i.repo.ReadCommand(ctx, app.ID(in.AppID), command.ID(in.ComID))
	if err != nil {
		return
	}

	c.SetName(command.Name(in.Name))
	err = i.repo.UpdateCommand(ctx, app.ID(in.AppID), c)

	return
}

func (i *Interactor) setIRData(ctx context.Context, in bdy.SetIRDataInput) (err error) {
	var c command.Command

	c, err = i.repo.ReadCommand(ctx, app.ID(in.AppID), command.ID(in.ComID))
	if err != nil {
		return
	}

	c.SetRawIRData(irdata.IRData(in.Data))
	err = i.repo.UpdateCommand(ctx, app.ID(in.AppID), c)
	return
}

// Delete
func (i *Interactor) deleteAppliance(ctx context.Context, in bdy.DeleteAppInput) (err error) {
	if _, err = i.repo.ReadApp(ctx, app.ID(in.AppID)); err != nil {
		return err
	}

	err = i.repo.DeleteApp(ctx, app.ID(in.AppID))
	if err != nil {
		return err
	}

	i.output.NotificateApplianceUpdate(
		ctx,
		bdy.UpdateNotifyOutput{
			AppID: in.AppID,
			Type:  bdy.UpdateTypeDelete,
		},
	)
	return err
}

func (i *Interactor) deleteCommand(ctx context.Context, in bdy.DeleteCommandInput) (err error) {
	var a *app.Appliance

	a, err = i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	err = a.RemoveCommand()
	if err != nil {
		return
	}

	err = i.repo.DeleteCommand(ctx, app.ID(in.AppID), command.ID(in.ComID))
	return
}
