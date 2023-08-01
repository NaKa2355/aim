package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *Handler) DeleteButton(ctx context.Context, req *aimv1.DeleteButtonRequest) (e *empty.Empty, err error) {
	var in bdy.DeleteButtonInput
	e = &empty.Empty{}

	in.RemoteID = req.RemoteId
	in.ButtonID = req.ButtonId
	err = h.i.DeleteButton(ctx, in)
	return
}
