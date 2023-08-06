package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
)

func (h *Handler) AddButton(ctx context.Context, req *aimv1.AddButtonRequest) (res *aimv1.AddButtonsResponse, err error) {
	res = &aimv1.AddButtonsResponse{}

	in := bdy.AddButtonInput{
		RemoteID: req.RemoteId,
		Name:     req.Name,
		Tag:      req.Tag,
	}

	b, err := h.i.AddButton(ctx, in)

	res.ButtonId = b.Button.ID
	return
}
