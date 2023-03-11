package interactor

import (
	"context"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

type Interactor struct {
	rep repository.Repository
}

var _ boundary.Boundary = &Interactor{}

func New(r repository.Repository) *Interactor {
	i := &Interactor{
		rep: r,
	}
	return i
}

func (i *Interactor) AddSwitch(ctx context.Context, d boundary.AddSwitch) (boundary.Appliance, error) {
	var a = boundary.Appliance{}
	name, err := appliance.NewName(d.Name)
	if err != nil {
		return a, err
	}

	devID, err := appliance.NewDeviceID(d.DeviceID)
	if err != nil {
		return a, err
	}

	s := appliance.NewSwitch(name, devID)
	app, err := i.rep.SaveApp(ctx, s)
	if err != nil {
		return a, err
	}

	a.Name = string(app.GetName())
	a.ID = string(app.GetID())
	a.ApplianceType = boundary.ApplianceType(app.GetType())
	a.DeviceID = string(app.GetDeviceID())
	return a, nil
}

func (i *Interactor) AddButton(ctx context.Context, name string, deviceID string) (boundary.Appliance, error) {
	var a = boundary.Appliance{}
	n, err := appliance.NewName(name)
	if err != nil {
		return a, err
	}

	d, err := appliance.NewDeviceID(deviceID)
	if err != nil {
		return a, err
	}

	b := appliance.NewButton(n, d)
	app, err := i.rep.SaveApp(ctx, b)
	if err != nil {
		return a, err
	}
	a.Name = string(app.GetName())
	a.ID = string(app.GetID())
	a.ApplianceType = boundary.ApplianceType(app.GetType())
	a.DeviceID = string(app.GetDeviceID())
	return a, nil
}

func (i *Interactor) AddThermostat(ctx context.Context, name string, deviceID string,
	scale float64,
	minimumHeatingTemp int,
	maximumHeatingTemp int,
	minimumCoolingTemp int,
	maximumCoolingTemp int) (boundary.Appliance, error) {

	var a = boundary.Appliance{}
	tOpt, err := appliance.NewThermostatOpt(scale, minimumHeatingTemp, maximumHeatingTemp, minimumCoolingTemp, maximumCoolingTemp)
	if err != nil {
		return a, err
	}

	n, err := appliance.NewName(name)
	if err != nil {
		return a, err
	}

	d, err := appliance.NewDeviceID(deviceID)
	if err != nil {
		return a, err
	}

	t := appliance.NewThermostat(n, d, tOpt)

	app, err := i.rep.SaveApp(ctx, t)
	if err != nil {
		return a, err
	}

	a.Name = string(app.GetName())
	a.ID = string(app.GetID())
	a.ApplianceType = boundary.ApplianceType(app.GetType())
	a.DeviceID = string(app.GetDeviceID())
	return a, nil
}

func (i *Interactor) AddCustom(ctx context.Context, name string, deviceID string) (boundary.Appliance, error) {
	var a = boundary.Appliance{}
	n, err := appliance.NewName(name)
	if err != nil {
		return a, err
	}

	d, err := appliance.NewDeviceID(deviceID)
	if err != nil {
		return a, err
	}

	c := appliance.NewCustom(n, d)
	app, err := i.rep.SaveApp(ctx, c)
	if err != nil {
		return a, err
	}

	a.Name = string(app.GetName())
	a.ID = string(app.GetID())
	a.ApplianceType = boundary.ApplianceType(app.GetType())
	a.DeviceID = string(app.GetDeviceID())
	return a, nil
}

func (i *Interactor) GetAppliances(ctx context.Context) ([]boundary.Appliance, error) {
	var boundaryApps []boundary.Appliance
	apps, err := i.rep.GetAppsList(ctx)
	if err != nil {
		return boundaryApps, err
	}
	boundaryApps = make([]boundary.Appliance, len(apps))
	for i, app := range apps {
		boundaryApps[i].Name = string(app.GetName())
		boundaryApps[i].ID = string(app.GetID())
		boundaryApps[i].DeviceID = string(app.GetDeviceID())
		boundaryApps[i].ApplianceType = boundary.ApplianceType(app.GetType())
	}
	return boundaryApps, nil
}

func (i *Interactor) RenameAppliance(ctx context.Context, appID string, name string) error {
	id, err := appliance.NewID(appID)
	if err != nil {
		return err
	}

	app, err := i.rep.GetApp(ctx, id)
	if err != nil {
		return err
	}

	n, err := appliance.NewName(name)
	if err != nil {
		return err
	}

	app = app.ChangeName(n)

	_, err = i.rep.SaveApp(ctx, app)
	return err
}

func (i *Interactor) ChangeIRDevice(ctx context.Context, appID string, irDevID string) error {
	id, err := appliance.NewID(appID)
	if err != nil {
		return err
	}

	app, err := i.rep.GetApp(ctx, id)
	if err != nil {
		return err
	}

	d, err := appliance.NewDeviceID(irDevID)
	if err != nil {
		return err
	}

	app = app.ChangeDeviceID(d)
	_, err = i.rep.SaveApp(ctx, app)
	return err
}

func (i *Interactor) DeleteAppliance(ctx context.Context, appID string) error {
	id, err := appliance.NewID(appID)
	if err != nil {
		return err
	}
	return i.rep.RemoveApp(ctx, id)
}

func (i *Interactor) GetCommands(ctx context.Context, appID string) ([]boundary.Command, error) {
	var bCom []boundary.Command

	id, err := appliance.NewID(appID)
	if err != nil {
		return bCom, err
	}

	coms, err := i.rep.GetCommands(ctx, id)
	if err != nil {
		return bCom, err
	}
	bCom = make([]boundary.Command, len(coms))
	for i, com := range coms {
		bCom[i].ID = string(com.GetID())
		bCom[i].Name = string(com.GetName())
	}
	return bCom, nil
}

func (i *Interactor) RenameCommand(ctx context.Context, appID string, comID string, name string) error {
	aID, err := appliance.NewID(appID)
	if err != nil {
		return err
	}

	cID, err := command.NewID(comID)
	if err != nil {
		return err
	}

	cName, err := command.NewName(name)
	if err != nil {
		return err
	}

	app, err := i.rep.GetApp(ctx, aID)
	if err != nil {
		return err
	}

	if app.GetType() != appliance.AppTypeCustom {
		return fmt.Errorf("this appliance is not support renaming")
	}

	com, err := i.rep.GetCommand(ctx, cID)
	if err != nil {
		return err
	}
	com = com.ChangeName(cName)

	_, err = i.rep.SaveCommand(ctx, aID, com)
	return err
}

func (i *Interactor) AddCommand(ctx context.Context, appID string, name string) error {
	id, err := appliance.NewID(appID)
	if err != nil {
		return err
	}

	n, err := command.NewName(name)
	if err != nil {
		return err
	}
	_, err = i.rep.SaveCommand(ctx, id, command.New("", n, nil))
	return err
}

func (i *Interactor) RemoveCommand(ctx context.Context, appID string, comID string) error {
	aID, err := appliance.NewID(appID)
	if err != nil {
		return err
	}

	cID, err := command.NewID(comID)
	if err != nil {
		return err
	}

	app, err := i.rep.GetApp(ctx, aID)
	if err != nil {
		return err
	}

	if app.GetType() != appliance.AppTypeCustom {
		return fmt.Errorf("this appliance is not support renaming")
	}

	return i.rep.RemoveCommand(ctx, cID)
}

func (i *Interactor) GetRawIRData(ctx context.Context, comID string) (boundary.RawIrData, error) {
	var irdata boundary.RawIrData
	id, err := command.NewID(comID)
	if err != nil {
		return irdata, err
	}

	com, err := i.rep.GetCommand(ctx, id)
	if err != nil {
		return irdata, err
	}

	irdata = boundary.RawIrData(com.GetRawIRData())
	return irdata, nil
}

func (i *Interactor) SetRawIRData(ctx context.Context, comID string, rawIRData boundary.RawIrData) error {
	id, err := command.NewID(comID)
	if err != nil {
		return err
	}
	return i.rep.SetRawIRData(ctx, id, irdata.RawIRData(rawIRData))
}
