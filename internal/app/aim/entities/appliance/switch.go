package appliance

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

func NewSwitch(name Name, deviceID DeviceID) (Appliance, error) {
	var a *ApplianceData
	a, err := NewAppliance(name, AppTypeSwitch, deviceID, "")
	if err != nil {
		return a, err
	}

	a.commands = append(a.commands, command.New("on"))
	a.commands = append(a.commands, command.New("off"))
	return a, nil
}
