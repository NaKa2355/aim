package interactor

import (
	"context"
	"fmt"

	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	com "github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

type Interactor struct {
	repo repository.Repository
}

var _ bdy.Boundary = &Interactor{}

func New(r repository.Repository) *Interactor {
	i := &Interactor{
		repo: r,
	}
	return i
}

func (i *Interactor) AddSwitch(ctx context.Context, d bdy.AddSwitch) (bdy.Appliance, error) {
	var r bdy.Appliance
	name, err := app.NewName(d.Name)
	if err != nil {
		return r, err
	}

	devID, err := app.NewDeviceID(d.DeviceID)
	if err != nil {
		return r, err
	}

	s := app.NewSwitch(name, devID)
	app, err := i.repo.SaveApp(ctx, s)
	if err != nil {
		return r, err
	}

	return convertApp(app), nil
}

func (i *Interactor) AddButton(ctx context.Context, d bdy.AddButton) (bdy.Appliance, error) {
	var r bdy.Appliance

	name, err := app.NewName(d.Name)
	if err != nil {
		return r, err
	}

	devID, err := app.NewDeviceID(d.DeviceID)
	if err != nil {
		return r, err
	}

	b := app.NewButton(name, devID)

	app, err := i.repo.SaveApp(ctx, b)
	if err != nil {
		return r, err
	}

	return convertApp(app), nil
}

func (i *Interactor) AddThermostat(ctx context.Context, d bdy.AddThermostat) (bdy.Appliance, error) {
	var r bdy.Appliance
	tOpt, err := app.NewThermostatOpt(d.Scale, d.MinimumHeatingTemp, d.MaximumHeatingTemp, d.MinimumCoolingTemp, d.MaximumCoolingTemp)
	if err != nil {
		return r, err
	}

	name, err := app.NewName(d.Name)
	if err != nil {
		return r, err
	}

	devID, err := app.NewDeviceID(d.DeviceID)
	if err != nil {
		return r, err
	}

	t := app.NewThermostat(name, devID, tOpt)

	app, err := i.repo.SaveApp(ctx, t)
	if err != nil {
		return r, err
	}

	return convertApp(app), nil
}

func (i *Interactor) AddCustom(ctx context.Context, d bdy.AddCustom) (bdy.Appliance, error) {
	var r bdy.Appliance
	name, err := app.NewName(d.Name)
	if err != nil {
		return r, err
	}

	devID, err := app.NewDeviceID(d.DeviceID)
	if err != nil {
		return r, err
	}

	c := app.NewCustom(name, devID)
	app, err := i.repo.SaveApp(ctx, c)
	if err != nil {
		return r, err
	}

	return convertApp(app), nil
}

func (i *Interactor) GetAppliances(ctx context.Context) ([]bdy.Appliance, error) {
	var boundaryApps []bdy.Appliance
	apps, err := i.repo.GetAppsList(ctx)
	if err != nil {
		return boundaryApps, err
	}
	boundaryApps = make([]bdy.Appliance, len(apps))
	for i, app := range apps {
		boundaryApps[i] = convertApp(app)
	}
	return boundaryApps, nil
}

func (i *Interactor) RenameAppliance(ctx context.Context, d bdy.RenameApp) error {
	id, err := app.NewID(d.AppID)
	if err != nil {
		return err
	}

	a, err := i.repo.GetApp(ctx, id)
	if err != nil {
		return err
	}

	name, err := app.NewName(d.Name)
	if err != nil {
		return err
	}

	a = a.ChangeName(name)

	_, err = i.repo.SaveApp(ctx, a)
	return err
}

func (i *Interactor) ChangeIRDevice(ctx context.Context, d bdy.ChangeIRDev) error {
	id, err := app.NewID(d.AppID)
	if err != nil {
		return err
	}

	a, err := i.repo.GetApp(ctx, id)
	if err != nil {
		return err
	}

	devID, err := app.NewDeviceID(d.DeviceID)
	if err != nil {
		return err
	}

	a = a.ChangeDeviceID(devID)
	_, err = i.repo.SaveApp(ctx, a)
	return err
}

func (i *Interactor) DeleteAppliance(ctx context.Context, d bdy.DeleteApp) error {
	id, err := app.NewID(d.AppID)
	if err != nil {
		return err
	}
	return i.repo.RemoveApp(ctx, id)
}

func (i *Interactor) GetCommands(ctx context.Context, d bdy.GetCommands) ([]bdy.Command, error) {
	var bCom []bdy.Command

	id, err := app.NewID(d.AppID)
	if err != nil {
		return bCom, err
	}

	coms, err := i.repo.GetCommands(ctx, id)
	if err != nil {
		return bCom, err
	}
	bCom = make([]bdy.Command, len(coms))
	for i, c := range coms {
		bCom[i].ID = string(c.GetID())
		bCom[i].Name = string(c.GetName())
	}
	return bCom, nil
}

func (i *Interactor) RenameCommand(ctx context.Context, d bdy.RenameCommand) error {
	appID, err := app.NewID(d.AppID)
	if err != nil {
		return err
	}

	comID, err := com.NewID(d.ComID)
	if err != nil {
		return err
	}

	name, err := com.NewName(d.Name)
	if err != nil {
		return err
	}

	a, err := i.repo.GetApp(ctx, appID)
	if err != nil {
		return err
	}

	if a.GetType() != app.TypeCustom {
		return fmt.Errorf("this appliance is not support renaming command(s)")
	}

	c, err := i.repo.GetCommand(ctx, comID)
	if err != nil {
		return err
	}
	c = c.ChangeName(name)

	_, err = i.repo.SaveCommand(ctx, appID, c)
	return err
}

func (i *Interactor) AddCommand(ctx context.Context, d bdy.AddCommand) error {
	id, err := app.NewID(d.AppID)
	if err != nil {
		return err
	}

	name, err := com.NewName(d.Name)
	if err != nil {
		return err
	}
	_, err = i.repo.SaveCommand(ctx, id, com.New("", name, nil))
	return err
}

func (i *Interactor) RemoveCommand(ctx context.Context, d bdy.RemoveCommand) error {
	aID, err := app.NewID(d.AppID)
	if err != nil {
		return err
	}

	cID, err := com.NewID(d.ComID)
	if err != nil {
		return err
	}

	a, err := i.repo.GetApp(ctx, aID)
	if err != nil {
		return err
	}

	if a.GetType() != app.TypeCustom {
		return fmt.Errorf("this appliance is not support removing command(s)")
	}

	return i.repo.RemoveCommand(ctx, cID)
}

func (i *Interactor) GetRawIRData(ctx context.Context, d bdy.GetRawIRData) (bdy.RawIRData, error) {
	var irdata bdy.RawIRData
	id, err := com.NewID(d.ComID)
	if err != nil {
		return irdata, err
	}

	c, err := i.repo.GetCommand(ctx, id)
	if err != nil {
		return irdata, err
	}

	irdata = bdy.RawIRData(c.GetRawIRData())
	return irdata, nil
}

func (i *Interactor) SetRawIRData(ctx context.Context, d bdy.SetRawIRData) error {
	id, err := com.NewID(d.ComID)
	if err != nil {
		return err
	}
	return i.repo.SetRawIRData(ctx, id, irdata.RawIRData(d.Data))
}
