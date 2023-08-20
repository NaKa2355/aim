package boundary

import "context"

type UpdateType int

const (
	UpdateTypeAdd UpdateType = iota
	UpdateTypeDelete
	UpdateTypeUpdate
)

type UpdateNotifyOutput struct {
	Remote Remote
	Type   UpdateType
}

type RemoteUpdateNotifier interface {
	NotificateRemoteUpdate(ctx context.Context, o UpdateNotifyOutput)
}
