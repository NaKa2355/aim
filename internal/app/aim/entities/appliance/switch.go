package appliance

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

func NewSwitch(name Name, deviceID DeviceID) Appliance {
	return NewAppliance("", name, AppTypeSwitch, deviceID, "", []command.Command{
		command.New("", "on", nil),
		command.New("", "off", nil),
	})
}
