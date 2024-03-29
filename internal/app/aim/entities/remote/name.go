package remote

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Name string

func NewName(name string) (Name, error) {
	err := validation.Validate(
		name,
		validation.Required,
		validation.RuneLength(0, 25),
	)
	if err != nil {
		return Name(""), entities.NewError(
			entities.CodeInvaildInput,
			fmt.Errorf("validation error at name: %w", err),
		)
	}
	return Name(name), nil
}
