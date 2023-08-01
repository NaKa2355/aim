package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *Handler) EditRemote(ctx context.Context, req *aimv1.EditRemoteRequest) (e *empty.Empty, err error) {
	var in bdy.EditRemoteInput
	e = &empty.Empty{}

	in.RemoteID = req.RemoteId
	in.Name = req.Name
	in.DeviceID = req.DeviceId
	err = h.i.EditRemote(ctx, in)
	return
}
