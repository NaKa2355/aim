package appliance

import "github.com/NaKa2355/aim/internal/app/aim/entities/command"

func NewButton(name Name, deviceID DeviceID) Appliance {
	a := NewAppliance("", name, AppTypeButton, deviceID, "", []command.Command{command.New("", "push", nil)})
	return a
}
