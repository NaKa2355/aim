package command

import "github.com/NaKa2355/aim/internal/app/aim/entities/irdata"

type Command struct {
	id     string
	name   string
	irdata irdata.RawIRData
}

func NewWithID(id string, name string) *Command {
	c := &Command{
		id:   id,
		name: name,
	}
	return c
}

func New(name string) *Command {
	c := &Command{
		name: name,
	}
	return c
}

func (c *Command) SetIRData(irdata irdata.RawIRData) {
	c.irdata = irdata
}

func (c *Command) GetName() string {
	return c.name
}

func (c *Command) GetRawIRData() irdata.RawIRData {
	return c.irdata
}

func (c *Command) GetID() string {
	return c.id
}

func (c *Command) SetID(id string) error {
	c.id = id
	return nil
}
