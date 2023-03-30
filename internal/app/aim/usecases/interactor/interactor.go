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
	var a app.Appliance
	switch in := _in.(type) {
	case bdy.AddCustomInput:
		a = app.NewCustom("", app.Name(in.Name), app.DeviceID(in.DeviceID))
	case bdy.AddButtonInput:
		a = app.NewButton("", app.Name(in.Name), app.DeviceID(in.DeviceID))
	case bdy.AddToggleInput:
		a = app.NewToggle("", app.Name(in.Name), app.DeviceID(in.DeviceID))
	case bdy.AddThermostatInput:
		a, err = app.NewThermostat("",
			app.Name(in.Name),
			app.DeviceID(in.DeviceID),
			in.Scale,
			in.MinimumHeatingTemp,
			in.MaximumHeatingTemp,
			in.MinimumCoolingTemp,
			in.MaximumCoolingTemp,
		)
		if err != nil {
			return
		}
	default:
		return out, errors.New("invaild input")
	}

	a, err = i.repo.CreateAppliance(ctx, a)
	out.ID = string(a.GetID())
	return
}

/*
func (i *Interactor) addCustom(ctx context.Context, in bdy.AddCustomInput) (bdy.AddAppOutput, error) {
	var out = bdy.AddAppOutput{}
	var err error = nil
	c := app.NewCustom("", app.Name(in.Name), app.DeviceID(in.DeviceID))
	c, err = i.repo.CreateCustom(ctx, c)
	if err != nil {
		return out, err
	}

	out.ID = string(c.ID)
	return out, nil
}

func (i *Interactor) addToggle(ctx context.Context, in bdy.AddToggleInput) (bdy.AddAppOutput, error) {
	var out = bdy.AddAppOutput{}

	t := app.NewToggle("", app.Name(in.Name), app.DeviceID(in.DeviceID))

	t, err := i.repo.CreateToggle(ctx, t)
	if err != nil {
		return out, err
	}

	out.ID = string(t.ID)
	return out, err
}

func (i *Interactor) addButton(ctx context.Context, in bdy.AddButtonInput) (bdy.AddAppOutput, error) {
	var out = bdy.AddAppOutput{}

	b := app.NewButton("", app.Name(in.Name), app.DeviceID(in.DeviceID))

	b, err := i.repo.CreateButton(ctx, b)
	if err != nil {
		return out, err
	}

	out.ID = string(b.ID)
	return out, err
}

func (i *Interactor) addThermostat(ctx context.Context, in bdy.AddThermostatInput) (out bdy.AddAppOutput, err error) {
	var t app.Thermostat

	t, err = app.NewThermostat("",
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
*/

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

/*
// Read
func (i *Interactor) getCustom(ctx context.Context, in bdy.GetAppInput) (out bdy.GetCustomOutput, err error) {
	var c app.Custom

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
	var t app.Toggle

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
	var b app.Button

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
	var t app.Thermostat

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
*/

func (i *Interactor) getAppliances(ctx context.Context) (out bdy.GetAppliancesOutput, err error) {
	var apps []app.Appliance

	apps, err = i.repo.ReadApps(ctx)
	if err != nil {
		return
	}

	out.Apps = make([]interface{}, len(apps))
	for i, _a := range apps {
		switch a := _a.(type) {
		case app.Custom:
			out.Apps[i] = bdy.Custom{
				ID:       string(a.ID),
				Name:     string(a.Name),
				DeviceID: string(a.DeviceID),
			}
		case app.Button:
			out.Apps[i] = bdy.Button{
				ID:       string(a.ID),
				Name:     string(a.Name),
				DeviceID: string(a.DeviceID),
			}
		case app.Toggle:
			out.Apps[i] = bdy.Toggle{
				ID:       string(a.ID),
				Name:     string(a.Name),
				DeviceID: string(a.DeviceID),
			}
		case app.Thermostat:
			out.Apps[i] = bdy.Thermostat{
				ID:                 string(a.ID),
				Name:               string(a.Name),
				DeviceID:           string(a.DeviceID),
				Scale:              a.Scale,
				MaximumHeatingTemp: a.MaximumHeatingTemp,
				MinimumHeatingTemp: a.MinimumHeatingTemp,
				MaximumCoolingTemp: a.MaximumCoolingTemp,
				MinimumCoolingTemp: a.MinimumCoolingTemp,
			}
		}
	}
	return
}

func (i *Interactor) getCommands(ctx context.Context, in bdy.GetCommandsInput) (out bdy.GetCommandsOutput, err error) {
	var coms []command.Command
	coms, err = i.repo.ReadCommands(ctx, app.ID(in.AppID))
	if err != nil {
		return
	}

	out.Commands = make([]bdy.Command, len(coms))
	for i, c := range coms {
		out.Commands[i].ID = string(c.ID)
		out.Commands[i].Name = string(c.Name)
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
