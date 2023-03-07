package appliance

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Switch struct {
	*ApplianceData
}

func NewSwitch(name string, deviceID string) (*Switch, error) {
	var s *Switch
	a, err := NewAppliance(name, AppTypeSwitch, deviceID)
	if err != nil {
		return s, err
	}
	s = &Switch{
		ApplianceData: a,
	}
	s.commands = append(s.commands, command.New("on"))
	s.commands = append(s.commands, command.New("off"))
	return s, nil
}
