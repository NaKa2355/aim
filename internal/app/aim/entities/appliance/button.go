package appliance

import "github.com/NaKa2355/aim/internal/app/aim/entities/command"

type Button struct {
	*ApplianceData
}

func NewButton(name string, deviceID string) (*Button, error) {
	var b *Button
	a, err := NewAppliance(name, AppTypeButton, deviceID)
	if err != nil {
		return b, err
	}
	b = &Button{
		ApplianceData: a,
	}
	b.commands = append(b.commands, command.New("push"))
	return b, nil
}
