package command

import "github.com/NaKa2355/aim/internal/app/aim/entities/irdata"

type Name string

func NewName(name string) Name {
	return Name(name)
}

type ID string

func NewID(id string) ID {
	return ID(id)
}

type Command struct {
	ID     ID
	Name   Name
	IRData irdata.IRData
}

func New(id ID, name Name, irdata irdata.IRData) Command {
	return Command{
		ID:     id,
		Name:   name,
		IRData: irdata,
	}
}

func (c *Command) GetID() ID {
	return c.ID
}

func (c *Command) SetID(id ID) {
	c.ID = id
}

func (c *Command) GetName() Name {
	return c.Name
}

func (c *Command) SetName(name Name) {
	c.Name = name
}

func (c *Command) GetRawIRData() irdata.IRData {
	return c.IRData
}

func (c *Command) SetRawIRData(irdata irdata.IRData) {
	c.IRData = irdata
}
