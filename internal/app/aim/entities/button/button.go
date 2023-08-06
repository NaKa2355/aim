package button

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

type Tag string

func NewTag(tag string) Tag {
	return Tag(tag)
}

type Button struct {
	ID     ID
	Name   Name
	Tag    Tag
	IRData irdata.IRData
}

func New(name Name, tag Tag, irdata irdata.IRData) *Button {
	return &Button{
		Name:   name,
		Tag:    tag,
		IRData: irdata,
	}
}

func (b *Button) GetID() ID {
	return b.ID
}

func (b *Button) SetID(id ID) {
	b.ID = id
}

func (b *Button) GetName() Name {
	return b.Name
}

func (b *Button) SetName(name Name) {
	b.Name = name
}

func (b *Button) GetTag() Tag {
	return b.Tag
}

func (b *Button) SetTag(tag Tag) {
	b.Tag = tag
}

func (b *Button) GetRawIRData() irdata.IRData {
	return b.IRData
}

func (b *Button) SetRawIRData(irdata irdata.IRData) {
	b.IRData = irdata
}
