package boundary

import "context"

type GetRemotesOutput struct {
	Remotes []Remote
}

type RemotesGetter interface {
	GetRemotes(ctx context.Context) (GetRemotesOutput, error)
}
