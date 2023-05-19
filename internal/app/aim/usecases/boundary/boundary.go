package boundary

import (
	"context"
)

type RemoteUpdateNotifier interface {
	NotificateRemoteUpdate(ctx context.Context, o UpdateNotifyOutput)
}

type RemoteAdder interface {
	AddRemote(ctx context.Context, i AddRemoteInput) (AddRemoteOutput, error)
}

type ButtonAdder interface {
	AddButton(ctx context.Context, i AddButtonInput) error
}

type RemotesGetter interface {
	GetRemotes(ctx context.Context) (GetRemotesOutput, error)
}

type RemoteGetter interface {
	GetRemote(ctx context.Context, i GetRemoteInput) (GetRemoteOutput, error)
}

type ButtonsGetter interface {
	GetButtons(ctx context.Context, i GetButtonsInput) (GetButtonsOutput, error)
}

type IRDataGetter interface {
	GetIRData(ctx context.Context, i GetIRDataInput) (GetIRDataOutput, error)
}

type RemoteEditor interface {
	EditRemote(ctx context.Context, i EditRemoteInput) error
}

type ButtonEditor interface {
	EditButton(ctx context.Context, i EditButtonInput) error
}

type IRDataSetter interface {
	SetIRData(ctx context.Context, i SetIRDataInput) error
}

type RemoteDeleter interface {
	DeleteRemote(ctx context.Context, i DeleteRemoteInput) error
}

type ButtonDeleter interface {
	DeleteButton(ctx context.Context, i DeleteButtonInput) error
}
