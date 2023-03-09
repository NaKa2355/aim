package command

import "github.com/NaKa2355/aim/internal/app/aim/entities/irdata"

type Name string

func NewName(name string) (Name, error) {
	return Name(name), nil
}

type ID string

func NewID(id string) (ID, error) {
	return ID(id), nil
}

type CommandData struct {
	id     ID
	name   Name
	irdata irdata.RawIRData
}

type Command interface {
	GetID() ID
	GetName() Name
	GetRawIRData() irdata.RawIRData
	SetRawIRData(irdata.RawIRData)
}

func New(id ID, name Name, irdata irdata.RawIRData) Command {
	c := &CommandData{
		id:     id,
		name:   name,
		irdata: irdata,
	}
	return c
}

func (c *CommandData) SetRawIRData(irdata irdata.RawIRData) {
	c.irdata = irdata
}

func (c *CommandData) GetID() ID {
	return c.id
}

func (c *CommandData) GetName() Name {
	return c.name
}

func (c *CommandData) GetRawIRData() irdata.RawIRData {
	return c.irdata
}
