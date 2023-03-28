package interactor

import (
	"context"

	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/custom"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/thermostat"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/toggle"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (i *Interactor) addCustom(ctx context.Context, in bdy.AddCustomInput) (bdy.AddAppOutput, error) {
	var out = bdy.AddAppOutput{}
	var err error = nil
	c := custom.New("", app.Name(in.Name), app.DeviceID(in.DeviceID))
	c, err = i.repo.CreateCustom(ctx, c)
	if err != nil {
		return out, err
	}

	out.ID = string(c.ID)
	return out, nil
}

func (i *Interactor) addToggle(ctx context.Context, in bdy.AddToggleInput) (bdy.AddAppOutput, error) {
	var out = bdy.AddAppOutput{}

	t := toggle.New("", app.Name(in.Name), app.DeviceID(in.DeviceID))

	t, err := i.repo.CreateToggle(ctx, t)
	if err != nil {
		return out, err
	}

	out.ID = string(t.ID)
	return out, err
}

func (i *Interactor) addButton(ctx context.Context, in bdy.AddButtonInput) (bdy.AddAppOutput, error) {
	var out = bdy.AddAppOutput{}

	b := button.New("", app.Name(in.Name), app.DeviceID(in.DeviceID))

	b, err := i.repo.CreateButton(ctx, b)
	if err != nil {
		return out, err
	}

	out.ID = string(b.ID)
	return out, err
}

func (i *Interactor) addThermostat(ctx context.Context, in bdy.AddThermostatInput) (out bdy.AddAppOutput, err error) {
	var t thermostat.Thermostat

	t, err = thermostat.New("",
		app.Name(in.Name),
		app.DeviceID(in.DeviceID),
		in.Scale,
		in.MinimumHeatingTemp, in.MaximumHeatingTemp,
		in.MinimumCoolingTemp, in.MaximumCoolingTemp,
	)
	if err != nil {
		return
	}

	t, err = i.repo.CreateThermostat(ctx, t)
	if err != nil {
		return
	}

	out.ID = string(t.ID)
	return
}

func (i *Interactor) addCommand(ctx context.Context, in bdy.AddCommandInput) (err error) {
	var a app.Appliance
	var com command.Command

	a, err = i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	err = a.AddCommand()
	if err != nil {
		return
	}

	com = command.New("", command.Name(in.Name), irdata.IRData{})
	_, err = i.repo.CreateCommand(ctx, app.ID(in.AppID), com)
	return
}

// Read
func (i *Interactor) getCustom(ctx context.Context, in bdy.GetAppInput) (out bdy.GetCustomOutput, err error) {
	var c custom.Custom

	c, err = i.repo.ReadCustom(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	out.ID = string(c.ID)
	out.Name = string(c.Name)
	out.DeviceID = string(c.DeviceID)
	out.Commands = convertComs(c.Commands)
	return
}

func (i *Interactor) getToggle(ctx context.Context, in bdy.GetAppInput) (out bdy.GetToggleOutput, err error) {
	var t toggle.Toggle

	t, err = i.repo.ReadToggle(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	out.ID = string(t.ID)
	out.Name = string(t.Name)
	out.DeviceID = string(t.DeviceID)
	out.Commands = convertComs(t.Commands)

	return
}

func (i *Interactor) getButton(ctx context.Context, in bdy.GetAppInput) (out bdy.GetButtonOutput, err error) {
	var b button.Button

	b, err = i.repo.ReadButton(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	out.ID = string(b.ID)
	out.Name = string(b.Name)
	out.DeviceID = string(b.DeviceID)
	out.Commands = convertComs(b.Commands)

	return
}

func (i *Interactor) getThermostat(ctx context.Context, in bdy.GetAppInput) (out bdy.GetThermostatOutput, err error) {
	var t thermostat.Thermostat

	t, err = i.repo.ReadThermostat(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	out.ID = string(t.ID)
	out.Name = string(t.Name)
	out.DeviceID = string(t.DeviceID)
	out.Commands = convertComs(t.Commands)
	out.Scale = t.Scale
	out.MaximumHeatingTemp = t.MaximumHeatingTemp
	out.MinimumHeatingTemp = t.MinimumHeatingTemp
	out.MaximumCoolingTemp = t.MaximumCoolingTemp
	out.MinimumCoolingTemp = t.MinimumCoolingTemp

	return
}

func (i *Interactor) getAppliances(ctx context.Context) (out bdy.GetAppliancesOutput, err error) {
	var apps []app.Appliance

	apps, err = i.repo.ReadApps(ctx)
	if err != nil {
		return
	}

	out.Apps = make([]bdy.Appliance, len(apps))
	for i, a := range apps {
		out.Apps[i].ID = string(a.ID)
		out.Apps[i].Name = string(a.Name)
		out.Apps[i].DeviceID = string(a.DeviceID)
		out.Apps[i].ApplianceType = bdy.ApplianceType(a.Type)
	}

	return
}

func (i *Interactor) getCommand(ctx context.Context, in bdy.GetCommandInput) (out bdy.GetCommandOutput, err error) {
	var com command.Command

	com, err = i.repo.ReadCommand(ctx, app.ID(in.AppID), command.ID(in.ComID))
	if err != nil {
		return
	}

	out.ID = string(com.ID)
	out.Name = string(com.Name)
	out.Data = bdy.IRData(com.IRData)

	return
}

// Update
func (i *Interactor) renameAppliance(ctx context.Context, in bdy.RenameAppInput) (err error) {
	var a app.Appliance

	a, err = i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	a.SetName(app.Name(in.Name))
	err = i.repo.UpdateApp(ctx, a)
	return
}

func (i *Interactor) changeIRDevice(ctx context.Context, in bdy.ChangeIRDevInput) (err error) {
	var a app.Appliance

	a, err = i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	a.SetDeviceID(app.DeviceID(in.DeviceID))
	err = i.repo.UpdateApp(ctx, a)

	return
}

func (i *Interactor) renameCommand(ctx context.Context, in bdy.RenameCommandInput) (err error) {
	var a app.Appliance
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
	err = i.repo.DeleteApp(ctx, app.ID(in.AppID))
	return
}

func (i *Interactor) deleteCommand(ctx context.Context, in bdy.DeleteCommandInput) (err error) {
	var a app.Appliance

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

func convertComs(coms []command.Command) []bdy.Command {
	in := make([]bdy.Command, len(coms))
	for i, c := range coms {
		in[i].ID = string(c.ID)
		in[i].Name = string(c.Name)
	}
	return in
}
