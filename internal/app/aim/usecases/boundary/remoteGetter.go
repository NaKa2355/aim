package boundary

import "context"

type GetRemoteOutput struct {
	Remote Remote
}

type GetRemoteInput struct {
	RemoteID string
}

type RemoteGetter interface {
	GetRemote(ctx context.Context, i GetRemoteInput) (GetRemoteOutput, error)
}
