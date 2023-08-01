package boundary

import "context"

type DeleteRemoteInput struct {
	RemoteID string
}

type RemoteDeleter interface {
	DeleteRemote(ctx context.Context, i DeleteRemoteInput) error
}
