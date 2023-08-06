package remote

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
)

type Tag string

type Remote struct {
	ID       ID
	Name     Name
	DeviceID DeviceID
	Tag      Tag
	Buttons  []*button.Button
}

func NewRemote(name string, deviceID string, tag string, buttons []*button.Button) (*Remote, error) {
	var remote *Remote
	n, err := NewName(name)
	if err != nil {
		return remote, err
	}

	d, err := NewDeviceID(deviceID)
	if err != nil {
		return remote, err
	}

	remote = &Remote{
		Name:     n,
		Tag:      Tag(tag),
		DeviceID: d,
		Buttons:  buttons,
	}
	return remote, err
}

func LoadAppliance(id ID, name Name, deviceID DeviceID, tag Tag) *Remote {
	return &Remote{
		ID:       id,
		Name:     name,
		DeviceID: deviceID,
		Tag:      tag,
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

func (a *Remote) SetTag(tag string) error {
	a.Tag = Tag(tag)
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
