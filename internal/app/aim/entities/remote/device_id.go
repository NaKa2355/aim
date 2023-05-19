package remote

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	validation "github.com/go-ozzo/ozzo-validation"
)

type DeviceID string

func NewDeviceID(name string) (DeviceID, error) {
	err := validation.Validate(name,
		validation.Required,
	)
	if err != nil {
		return DeviceID(""), entities.NewError(
			entities.CodeInvaildInput,
			fmt.Errorf("validation error at device_id: %w", err),
		)
	}
	return DeviceID(name), nil
}
