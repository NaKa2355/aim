package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *Handler) DeleteRemote(ctx context.Context, req *aimv1.DeleteRemoteRequest) (e *empty.Empty, err error) {
	var in bdy.DeleteRemoteInput
	e = &empty.Empty{}

	in.RemoteID = req.RemoteId

	err = h.i.DeleteRemote(ctx, in)
	return
}
