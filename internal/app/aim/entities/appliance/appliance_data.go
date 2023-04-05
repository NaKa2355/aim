package appliance

import "github.com/NaKa2355/aim/internal/app/aim/entities/command"

type ApplianceData struct {
	ID       ID
	Name     Name
	DeviceID DeviceID
	Commands []command.Command
}

func NewApplianceData(name string, deviceID string, commands []command.Command) (a *ApplianceData, err error) {
	n, err := NewName(name)
	if err != nil {
		return a, err
	}

	d, err := NewDeviceID(deviceID)
	if err != nil {
		return a, err
	}

	return &ApplianceData{
		ID:       "",
		Name:     n,
		DeviceID: d,
		Commands: commands,
	}, nil
}

func (a *ApplianceData) GetID() ID {
	return a.ID
}

func (a *ApplianceData) GetName() Name {
	return a.Name
}

func (a *ApplianceData) GetDeviceID() DeviceID {
	return a.DeviceID
}

func (a *ApplianceData) GetCommands() []command.Command {
	return a.Commands
}

func (a *ApplianceData) SetID(id string) (err error) {
	a.ID, err = NewID(id)
	return
}

func (a *ApplianceData) SetName(name string) (err error) {
	a.Name, err = NewName(name)
	return
}

func (b *ApplianceData) SetDeviceID(deviceID string) (err error) {
	b.DeviceID, err = NewDeviceID(deviceID)
	return
}
