package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *Handler) NotificateRemoteUpdate(ctx context.Context, o bdy.UpdateNotifyOutput) {
	defer h.c.L.Unlock()
	notification := aimv1.UpdateNotification{}

	switch o.Type {
	case bdy.UpdateTypeAdd:
		notification.Notification = &aimv1.UpdateNotification_Add{
			Add: &aimv1.RemoteAdditionNotification{
				Remote: &aimv1.Remote{
					Id:       o.Remote.ID,
					Name:     o.Remote.Name,
					DeviceId: o.Remote.DeviceID,
					Tag:      o.Remote.Tag,
				},
			},
		}
	case bdy.UpdateTypeDelete:
		notification.Notification = &aimv1.UpdateNotification_Delete{
			Delete: &aimv1.RemoteDeletionNotification{
				RemoteId: o.Remote.ID,
			},
		}
	}

	h.c.L.Lock()
	h.notification = notification
	h.c.Broadcast()
}

func (h *Handler) NotifyRemoteUpdate(_ *empty.Empty, stream aimv1.AimService_NotifyUpdateServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case <-h.c.NotifyChan():
			h.c.L.Lock()
			err := stream.Send(&h.notification)
			h.c.L.Unlock()
			if err != nil {
				return err
			}
		}
	}
}
