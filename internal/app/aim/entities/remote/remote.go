package remote

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
)

type RemoteType int

const (
	TypeCustom RemoteType = iota
	TypeButton
	TypeToggle
	TypeThermostat
)

type Remote struct {
	ID       ID
	Name     Name
	DeviceID DeviceID
	Type     RemoteType
	Buttons  []*button.Button
	RemoteOperator
}

type RemoteOperator interface {
	ChangeButtonName() error
	AddButton() error
	RemoveButton() error
}

func NewAppliance(n string, d string, remoteType RemoteType, buttons []*button.Button, opr RemoteOperator) (*Remote, error) {
	var remote *Remote
	name, err := NewName(n)
	if err != nil {
		return remote, err
	}

	deviceID, err := NewDeviceID(d)
	if err != nil {
		return remote, err
	}

	remote = &Remote{
		Name:           name,
		Type:           remoteType,
		DeviceID:       deviceID,
		Buttons:        buttons,
		RemoteOperator: opr,
	}
	return remote, err
}

func LoadAppliance(id ID, name Name, deviceID DeviceID, remoteType RemoteType, opr RemoteOperator) *Remote {
	return &Remote{
		ID:             id,
		Name:           name,
		DeviceID:       deviceID,
		Type:           remoteType,
		RemoteOperator: opr,
	}
}

func (r *Remote) SetID(id string) error {
	_id, err := NewID(id)
	if err != nil {
		return err
	}
	r.ID = _id
	return nil
}

func (a *Remote) SetName(name string) error {
	_name, err := NewName(name)
	if err != nil {
		return err
	}
	a.Name = _name
	return nil
}

func (r *Remote) SetDeviceID(devID string) error {
	_id, err := NewDeviceID(devID)
	if err != nil {
		return err
	}
	r.DeviceID = _id
	return nil
}
