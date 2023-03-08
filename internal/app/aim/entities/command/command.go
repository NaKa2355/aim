package command

import "github.com/NaKa2355/aim/internal/app/aim/entities/irdata"

type CommandData struct {
	name   string
	irdata irdata.RawIRData
}

type Command interface {
	GetName() string
	GetRawIRData() irdata.RawIRData
	SetRawIRData(irdata.RawIRData)
}

func New(name string) Command {
	c := &CommandData{
		name: name,
	}
	return c
}

func (c *CommandData) SetRawIRData(irdata irdata.RawIRData) {
	c.irdata = irdata
}

func (c *CommandData) GetName() string {
	return c.name
}

func (c *CommandData) GetRawIRData() irdata.RawIRData {
	return c.irdata
}
