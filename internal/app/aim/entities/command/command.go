package command

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Name string

func NewName(name string) (Name, error) {
	err := validation.Validate(name,
		validation.Required,
		validation.Length(1, 20),
	)
	if err != nil {
		return Name(""), entities.NewError(
			entities.CodeInvaildInput,
			fmt.Errorf("validation error at name: %w", err),
		)
	}
	return Name(name), nil
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

func New(name Name, irdata irdata.IRData) *Command {
	return &Command{
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
