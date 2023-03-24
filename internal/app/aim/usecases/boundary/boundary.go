package boundary

import "context"

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
	SetRawIRData(ctx context.Context, err error)

	//Delete
	DeleteAppliance(ctx context.Context, err error)
	DeleteCommand(ctx context.Context, err error)
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
	GetCommand(ctx context.Context, i GetRawIRDataInput)

	//Update
	RenameAppliance(ctx context.Context, i RenameAppInput)
	ChangeIRDevice(ctx context.Context, i ChangeIRDevInput)
	RenameCommand(ctx context.Context, i RenameCommandInput)
	SetRawIRData(ctx context.Context, i SetRawIRDataInput)

	//Delete
	DeleteAppliance(ctx context.Context, i DeleteAppInput)
	DeleteCommand(ctx context.Context, i RemoveCommandInput)
}
