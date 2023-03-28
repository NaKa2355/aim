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

/*

type OutputBoundary interface {
	//Create
	AddCustom(ctx context.Context, o AddAppOutput, err error)
	AddToggle(ctx context.Context, o AddAppOutput, err error)
	AddButton(ctx context.Context, o AddAppOutput, err error)
	AddThermostat(ctx context.Context, o AddAppOutput, err error)
	AddCommand(ctx context.Context, err error)

	//Read
	GetCustom(ctx context.Context, o GetCustomOutput, err error)
	GetToggle(ctx context.Context, o GetToggleOutput, err error)
	GetButton(ctx context.Context, o GetButtonOutput, err error)
	GetThermostat(ctx context.Context, o GetThermostatOutput, err error)
	GetAppliances(ctx context.Context, o GetAppliancesOutput, err error)
	GetCommand(ctx context.Context, c GetCommandOutput, err error)

	//Update
	RenameAppliance(ctx context.Context, err error)
	ChangeIRDevice(ctx context.Context, err error)
	RenameCommand(ctx context.Context, err error)
	SetIRData(ctx context.Context, err error)

	//Delete
	DeleteAppliance(ctx context.Context, err error)
	DeleteCommand(ctx context.Context, err error)

	ChangeNotify(o ChangeNotifyOutput)
}

type InputBoundary interface {
	//Create
	AddCustom(ctx context.Context, i AddCustomInput)
	AddToggle(ctx context.Context, i AddToggleInput)
	AddButton(ctx context.Context, i AddButtonInput)
	AddThermostat(ctx context.Context, i AddThermostatInput)
	AddCommand(ctx context.Context, i AddCommandInput)

	//Read
	GetCustom(ctx context.Context, i GetAppInput)
	GetToggle(ctx context.Context, i GetAppInput)
	GetButton(ctx context.Context, i GetAppInput)
	GetThermostat(ctx context.Context, i GetAppInput)
	GetAppliances(ctx context.Context)
	GetCommand(ctx context.Context, i GetCommandInput)

	//Update
	RenameAppliance(ctx context.Context, i RenameAppInput)
	ChangeIRDevice(ctx context.Context, i ChangeIRDevInput)
	RenameCommand(ctx context.Context, i RenameCommandInput)
	SetIRData(ctx context.Context, i SetIRDataInput)

	//Delete
	DeleteAppliance(ctx context.Context, i DeleteAppInput)
	DeleteCommand(ctx context.Context, i DeleteCommandInput)
}
*/
