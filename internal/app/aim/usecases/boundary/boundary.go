package boundary

import (
	"context"
)

type ApplianceUpdateNotifier interface {
	NotificateApplianceUpdate(ctx context.Context, o UpdateNotifyOutput)
}

type ApplianceAdder interface {
	AddAppliance(ctx context.Context, i AddApplianceInput) (AddAppOutput, error)
}

type CommandAdder interface {
	AddCommand(ctx context.Context, i AddCommandInput) error
}

type AppliancesGetter interface {
	GetAppliances(ctx context.Context) (GetAppliancesOutput, error)
}

type CommandsGetter interface {
	GetCommands(ctx context.Context, i GetCommandsInput) (GetCommandsOutput, error)
}

type CommandGetter interface {
	GetCommand(ctx context.Context, i GetCommandInput) (GetCommandOutput, error)
}

type IRDataGetter interface {
	GetIRData(ctx context.Context, i GetIRDataInput) (IRData, error)
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
