package boundary

import "context"

type EditRemoteInput struct {
	RemoteID string
	Name     string
	DeviceID string
}

type RemoteEditor interface {
	EditRemote(ctx context.Context, i EditRemoteInput) error
}
