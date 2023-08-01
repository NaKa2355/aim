package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *Handler) AddButton(ctx context.Context, req *aimv1.AddButtonRequest) (e *empty.Empty, err error) {
	e = &empty.Empty{}

	in := bdy.AddButtonInput{
		RemoteID: req.RemoteId,
		Name:     req.Name,
	}
	err = h.i.AddButton(ctx, in)
	return
}
