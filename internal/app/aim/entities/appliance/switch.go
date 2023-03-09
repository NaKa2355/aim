package appliance

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

func NewSwitch(name Name, deviceID DeviceID) (Appliance, error) {
	a := NewAppliance(name, AppTypeSwitch, deviceID, "")
	a.commands = append(a.commands, command.New("", "on"))
	a.commands = append(a.commands, command.New("", "off"))
	return a, nil
}
