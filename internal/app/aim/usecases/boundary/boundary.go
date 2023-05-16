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

type ApplianceGetter interface {
	GetAppliance(ctx context.Context, i GetApplianceInput) (GetApplianceOutput, error)
}

type CommandsGetter interface {
	GetCommands(ctx context.Context, i GetCommandsInput) (GetCommandsOutput, error)
}

type IRDataGetter interface {
	GetIRData(ctx context.Context, i GetIRDataInput) (GetIRDataOutput, error)
}

type ApplianceEditor interface {
	EditAppliance(ctx context.Context, i EditApplianceInput) error
}

type CommandEditor interface {
	EditCommand(ctx context.Context, i EditCommandInput) error
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
