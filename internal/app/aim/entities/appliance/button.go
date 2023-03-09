package appliance

import "github.com/NaKa2355/aim/internal/app/aim/entities/command"

func NewButton(name Name, deviceID DeviceID) (Appliance, error) {
	var a *ApplianceData
	a, err := NewAppliance(name, AppTypeButton, deviceID, "")
	if err != nil {
		return a, err
	}
	a.commands = append(a.commands, command.New("", "push"))
	return a, nil
}
