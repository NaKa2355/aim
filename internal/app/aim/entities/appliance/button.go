package appliance

import "github.com/NaKa2355/aim/internal/app/aim/entities/command"

func NewButton(name Name, deviceID DeviceID) Appliance {
	return NewAppliance("", name, TypeButton, deviceID, "", []command.Command{command.New("", "push", nil)})
}
