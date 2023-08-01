package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *Handler) EditButton(ctx context.Context, req *aimv1.EditButtonRequest) (e *empty.Empty, err error) {
	var in bdy.EditButtonInput
	e = &empty.Empty{}

	in.RemoteID = req.RemoteId
	in.ButtonID = req.ButtonId
	in.Name = req.Name
	err = h.i.EditButton(ctx, in)
	return
}
