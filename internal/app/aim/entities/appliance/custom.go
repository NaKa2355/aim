package appliance

import "github.com/NaKa2355/aim/internal/app/aim/entities/command"

func NewCustom(name Name, deviceID DeviceID) Appliance {
	return NewAppliance("", name, TypeCustom, deviceID, "", make([]command.Command, 0))
}
