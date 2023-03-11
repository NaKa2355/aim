package boundary

import "context"

type Boundary interface {
	AddSwitch(ctx context.Context, d AddSwitch) (Appliance, error)
	AddButton(ctx context.Context, d AddButton) (Appliance, error)
	AddThermostat(ctx context.Context, d AddThermostat) (Appliance, error)
	AddCustom(ctx context.Context, d AddCustom) (Appliance, error)

	GetAppliances(ctx context.Context) ([]Appliance, error)
	RenameAppliance(ctx context.Context, d RenameApp) error
	ChangeIRDevice(ctx context.Context, d ChangeIRDev) error
	DeleteAppliance(ctx context.Context, d DeleteApp) error

	GetCommands(ctx context.Context, d GetCommands) ([]Command, error)
	RenameCommand(ctx context.Context, d RenameCommand) error
	AddCommand(ctx context.Context, d AddCommand) error
	RemoveCommand(ctx context.Context, d RemoveCommand) error
	GetRawIRData(ctx context.Context, d GetRawIRData) (RawIRData, error)
	SetRawIRData(ctx context.Context, d SetRawIRData) error
}
