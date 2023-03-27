package interactor

import (
	"context"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/custom"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/thermostat"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/toggle"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

type Interactor struct {
	repo   repository.Repository
	output bdy.OutputBoundary
}

var _ bdy.InputBoundary = &Interactor{}

func wrapErr(err error) error {
	if err == nil {
		return nil
	}
	var code bdy.ErrCode

	if entityErr, ok := err.(entities.Error); ok {
		switch entityErr.Code {
		case entities.CodeInvaildInput:
			code = bdy.CodeInvaildInput
		case entities.CodeInvaildOperation:
			code = bdy.CodeInvaildOperation
		}
	}

	if repoErr, ok := err.(repository.Error); ok {
		switch repoErr.Code {
		case repository.CodeInvaildInput:
			code = bdy.CodeInvaildInput
		case repository.CodeNotFound:
			code = bdy.CodeNotFound
		case repository.CodeDataBase:
			code = bdy.CodeDatabase
		}
	}
	return bdy.NewError(code, err)
}

func New(in repository.Repository, o bdy.OutputBoundary) *Interactor {
	i := &Interactor{
		repo:   in,
		output: o,
	}
	return i
}

func (i *Interactor) AddCustom(ctx context.Context, in bdy.AddCustomInput) {
	var out = bdy.AddAppOutput{}
	c := custom.New("", app.Name(in.Name), app.DeviceID(in.DeviceID))
	c, err := i.repo.CreateCustom(ctx, c)
	if err != nil {
		i.output.AddCustom(ctx, out, wrapErr(err))
		return
	}

	out.ID = string(c.ID)
	i.output.AddCustom(ctx, out, wrapErr(err))
}

func (i *Interactor) AddToggle(ctx context.Context, in bdy.AddToggleInput) {
	var out = bdy.AddAppOutput{}

	t := toggle.New("", app.Name(in.Name), app.DeviceID(in.DeviceID))
	t, err := i.repo.CreateToggle(ctx, t)
	if err != nil {
		i.output.AddToggle(ctx, out, wrapErr(err))
		return
	}

	out.ID = string(t.ID)
	i.output.AddToggle(ctx, out, wrapErr(err))
}

func (i *Interactor) AddButton(ctx context.Context, in bdy.AddButtonInput) {
	var out = bdy.AddAppOutput{}

	b := button.New("", app.Name(in.Name), app.DeviceID(in.DeviceID))
	b, err := i.repo.CreateButton(ctx, b)
	if err != nil {
		i.output.AddButton(ctx, out, wrapErr(err))
		return
	}

	out.ID = string(b.ID)
	i.output.AddButton(ctx, out, wrapErr(err))
}

func (i *Interactor) AddThermostat(ctx context.Context, in bdy.AddThermostatInput) {
	var out bdy.AddAppOutput
	t, err := thermostat.New("", app.Name(in.Name), app.DeviceID(in.DeviceID),
		in.Scale, in.MinimumHeatingTemp, in.MaximumHeatingTemp, in.MinimumCoolingTemp, in.MaximumCoolingTemp)
	if err != nil {
		i.output.AddThermostat(ctx, out, wrapErr(err))
		return
	}

	t, err = i.repo.CreateThermostat(ctx, t)
	if err != nil {
		i.output.AddThermostat(ctx, out, wrapErr(err))
		return
	}

	out.ID = string(t.ID)
	i.output.AddThermostat(ctx, out, wrapErr(err))
}

func (i *Interactor) AddCommand(ctx context.Context, in bdy.AddCommandInput) {
	a, err := i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		i.output.AddCommand(ctx, wrapErr(err))
		return
	}
	if err = a.AddCommand(); err != nil {
		fmt.Println("hello")
		i.output.AddCommand(ctx, wrapErr(err))
		return
	}
	com := command.New("", command.Name(in.Name), irdata.IRData{})
	_, err = i.repo.CreateCommand(ctx, app.ID(in.AppID), com)

	i.output.AddCommand(ctx, wrapErr(err))
}

// Read
func (i *Interactor) GetCustom(ctx context.Context, in bdy.GetAppInput) {
	var out bdy.GetCustomOutput
	c, err := i.repo.ReadCustom(ctx, app.ID(in.AppID))
	if err != nil {
		i.output.GetCustom(ctx, out, wrapErr(err))
		return
	}

	out = bdy.GetCustomOutput{
		ID:       string(c.ID),
		Name:     string(c.Name),
		DeviceID: string(c.DeviceID),
		Command:  convertComs(c.Commands),
	}

	i.output.GetCustom(ctx, out, wrapErr(err))
}

func (i *Interactor) GetToggle(ctx context.Context, in bdy.GetAppInput) {
	var out bdy.GetToggleOutput
	t, err := i.repo.ReadToggle(ctx, app.ID(in.AppID))
	if err != nil {
		i.output.GetToggle(ctx, out, wrapErr(err))
		return
	}

	out = bdy.GetToggleOutput{
		ID:       string(t.ID),
		Name:     string(t.Name),
		DeviceID: string(t.DeviceID),
		Command:  convertComs(t.Commands),
	}
	i.output.GetToggle(ctx, out, wrapErr(err))
}

func (i *Interactor) GetButton(ctx context.Context, in bdy.GetAppInput) {
	var out bdy.GetButtonOutput
	b, err := i.repo.ReadButton(ctx, app.ID(in.AppID))
	if err != nil {
		i.output.GetButton(ctx, out, wrapErr(err))
		return
	}
	out = bdy.GetButtonOutput{
		ID:       string(b.ID),
		Name:     string(b.Name),
		DeviceID: string(b.DeviceID),
		Command:  convertComs(b.Commands),
	}

	i.output.GetButton(ctx, out, wrapErr(err))
}

func (i *Interactor) GetThermostat(ctx context.Context, in bdy.GetAppInput) {
	var out bdy.GetThermostatOutput
	t, err := i.repo.ReadThermostat(ctx, app.ID(in.AppID))
	if err != nil {
		i.output.GetThermostat(ctx, out, wrapErr(err))
		return
	}

	out = bdy.GetThermostatOutput{
		ID:                 string(t.ID),
		Name:               string(t.Name),
		DeviceID:           string(t.DeviceID),
		Command:            convertComs(t.Commands),
		Scale:              t.Scale,
		MaximumHeatingTemp: t.MaximumHeatingTemp,
		MinimumHeatingTemp: t.MinimumHeatingTemp,
		MaximumCoolingTemp: t.MaximumCoolingTemp,
		MinimumCoolingTemp: t.MinimumCoolingTemp,
	}

	i.output.GetThermostat(ctx, out, wrapErr(err))
}

func (i *Interactor) GetAppliances(ctx context.Context) {
	var out = bdy.GetAppliancesOutput{}
	apps, err := i.repo.ReadApps(ctx)
	if err != nil {
		i.output.GetAppliances(ctx, out, wrapErr(err))
		return
	}
	out.Apps = make([]bdy.Appliance, len(apps))

	for i, a := range apps {
		out.Apps[i] = bdy.Appliance{
			ID:            string(a.ID),
			Name:          string(a.Name),
			DeviceID:      string(a.DeviceID),
			ApplianceType: bdy.ApplianceType(a.Type),
		}
	}
	i.output.GetAppliances(ctx, out, wrapErr(err))
}

func (i *Interactor) GetCommand(ctx context.Context, in bdy.GetRawIRDataInput) {
	var out bdy.GetCommandOutput
	com, err := i.repo.ReadCommand(ctx, app.ID(in.AppID), command.ID(in.ComID))
	if err != nil {
		i.output.GetCommand(ctx, out, wrapErr(err))
		return
	}

	out = bdy.GetCommandOutput{
		ID:   string(com.ID),
		Name: string(com.Name),
		Data: bdy.IRData(com.IRData),
	}

	i.output.GetCommand(ctx, out, wrapErr(err))
}

// Update
func (i *Interactor) RenameAppliance(ctx context.Context, in bdy.RenameAppInput) {
	a, err := i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		i.output.RenameAppliance(ctx, wrapErr(err))
		return
	}

	a.SetName(app.Name(in.Name))
	err = i.repo.UpdateApp(ctx, a)
	i.output.RenameAppliance(ctx, wrapErr(err))
}

func (i *Interactor) ChangeIRDevice(ctx context.Context, in bdy.ChangeIRDevInput) {
	a, err := i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		i.output.ChangeIRDevice(ctx, wrapErr(err))
		return
	}

	a.SetDeviceID(app.DeviceID(in.DeviceID))
	err = i.repo.UpdateApp(ctx, a)
	i.output.ChangeIRDevice(ctx, wrapErr(err))
}

func (i *Interactor) RenameCommand(ctx context.Context, in bdy.RenameCommandInput) {
	a, err := i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		i.output.RenameCommand(ctx, wrapErr(err))
		return
	}

	if err = a.ChangeCommandName(); err != nil {
		i.output.RenameCommand(ctx, wrapErr(err))
		return
	}

	c, err := i.repo.ReadCommand(ctx, app.ID(in.AppID), command.ID(in.ComID))
	if err != nil {
		i.output.RenameCommand(ctx, wrapErr(err))
		return
	}

	c.SetName(command.Name(in.Name))
	err = i.repo.UpdateCommand(ctx, app.ID(in.AppID), c)
	i.output.RenameCommand(ctx, wrapErr(err))
}

func (i *Interactor) SetRawIRData(ctx context.Context, in bdy.SetRawIRDataInput) {
	c, err := i.repo.ReadCommand(ctx, app.ID(in.AppID), command.ID(in.ComID))
	if err != nil {
		i.output.SetRawIRData(ctx, wrapErr(err))
		return
	}
	c.SetRawIRData(irdata.IRData(in.Data))
	err = i.repo.UpdateCommand(ctx, app.ID(in.AppID), c)
	i.output.SetRawIRData(ctx, wrapErr(err))
}

// Delete
func (i *Interactor) DeleteAppliance(ctx context.Context, in bdy.DeleteAppInput) {
	err := i.repo.DeleteApp(ctx, app.ID(in.AppID))
	i.output.DeleteAppliance(ctx, wrapErr(err))
}

func (i *Interactor) DeleteCommand(ctx context.Context, in bdy.DeleteCommandInput) {
	a, err := i.repo.ReadApp(ctx, app.ID(in.AppID))
	if err != nil {
		i.output.DeleteCommand(ctx, wrapErr(err))
		return
	}

	if err = a.RemoveCommand(); err != nil {
		i.output.DeleteCommand(ctx, wrapErr(err))
		return
	}

	err = i.repo.DeleteCommand(ctx, app.ID(in.AppID), command.ID(in.ComID))
	i.output.DeleteCommand(ctx, wrapErr(err))
}

func convertComs(coms []command.Command) []bdy.Command {
	in := make([]bdy.Command, len(coms))
	for i, c := range coms {
		in[i].ID = string(c.ID)
		in[i].Name = string(c.Name)
	}
	return in
}
