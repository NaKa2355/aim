package boundary

import "context"

type Boundary interface {
	AddSwitch(ctx context.Context, d AddSwitch) (Appliance, error)
	AddButton(ctx context.Context, name string, deviceID string) (Appliance, error)
	AddThermostat(ctx context.Context,
		name string, deviceID string,
		scale float64,
		maximumHeatingTemp int,
		minimumHeatingTemp int,
		maximumCoolingTemp int,
		minimumCoolingTemp int) (Appliance, error)
	AddCustom(ctx context.Context, name string, deviceID string) (Appliance, error)

	GetAppliances(ctx context.Context) ([]Appliance, error)
	RenameAppliance(ctx context.Context, appID string, name string) error
	ChangeIRDevice(ctx context.Context, appID string, irDevID string) error
	DeleteAppliance(ctx context.Context, appID string) error

	GetCommands(ctx context.Context, appID string) ([]Command, error)
	RenameCommand(ctx context.Context, appID string, comID string, name string) error
	AddCommand(ctx context.Context, appID string, name string) error
	RemoveCommand(ctx context.Context, appID string, comID string) error
	GetRawIRData(ctx context.Context, comID string) (RawIrData, error)
	SetRawIRData(ctx context.Context, comID string, rawIRData RawIrData) error
}
