package boundary

import (
	"context"
)

type CustomAdder interface {
	AddCustom(ctx context.Context, i AddCustomInput) (AddAppOutput, error)
}

type ToggleAdder interface {
	AddToggle(ctx context.Context, i AddToggleInput) (AddAppOutput, error)
}

type ButtonAdder interface {
	AddButton(ctx context.Context, i AddButtonInput) (AddAppOutput, error)
}

type ThermostatAdder interface {
	AddThermostat(ctx context.Context, i AddThermostatInput) (AddAppOutput, error)
}

type CommandAdder interface {
	AddCommand(ctx context.Context, i AddCommandInput) error
}

type CustomGetter interface {
	GetCustom(ctx context.Context, i GetAppInput) (GetCustomOutput, error)
}

type ToggleGetter interface {
	GetToggle(ctx context.Context, i GetAppInput) (GetToggleOutput, error)
}

type ButtonGetter interface {
	GetButton(ctx context.Context, i GetAppInput) (GetButtonOutput, error)
}

type ThermostatGetter interface {
	GetThermostat(ctx context.Context, i GetAppInput) (GetThermostatOutput, error)
}

type AppliancesGetter interface {
	GetAppliances(ctx context.Context) (GetAppliancesOutput, error)
}

type CommandGetter interface {
	GetCommand(ctx context.Context, i GetCommandInput) (GetCommandOutput, error)
}

type ApplianceRenamer interface {
	RenameAppliance(ctx context.Context, i RenameAppInput) error
}

type IRDeviceChanger interface {
	ChangeIRDevice(ctx context.Context, i ChangeIRDevInput) error
}

type CommandRenamer interface {
	RenameCommand(ctx context.Context, i RenameCommandInput) error
}

type IRDataSetter interface {
	SetIRData(ctx context.Context, i SetIRDataInput) error
}

type ApplianceDeleter interface {
	DeleteAppliance(ctx context.Context, i DeleteAppInput) error
}

type CommandDeleter interface {
	DeleteCommand(ctx context.Context, i DeleteCommandInput) error
}
