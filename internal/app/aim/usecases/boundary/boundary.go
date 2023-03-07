package boundary

type Boundary interface {
	AddSwitch(name string) (Appliance, error)
	AddButton(name string) (Appliance, error)
	AddThermostat(name string,
		scale float64,
		maximumHeatingTemp int,
		minimumHeatingTemp int,
		maximumCoolingTemp int,
		minimumCoolingTemp int) (Appliance, error)

	GetAppliances() ([]*Appliance, error)
	RenameAppliance(appID string, name string) (Appliance, error)
	ChangeDevice(appID string, devID string) (Appliance, error)
	DeleteAppliance(appID string) error
	GetCommands(appID string) ([]*Command, error)
	GetRawIrData(comID string) (RawIrData, error)
	UpdateRawIrData(comID string, rawIRData RawIrData) error
}
