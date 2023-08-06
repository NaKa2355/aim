package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
)

func (h *Handler) GetButtons(ctx context.Context, req *aimv1.GetButtonsRequest) (res *aimv1.GetButtonsResponse, err error) {
	var in = bdy.GetButtonsInput{}
	var out bdy.GetButtonsOutput
	res = &aimv1.GetButtonsResponse{}

	in.RemoteID = req.RemoteId

	out, err = h.i.GetButtons(ctx, in)
	if err != nil {
		return
	}

	res.Buttons = make([]*aimv1.Button, len(out.Buttons))
	for i, b := range out.Buttons {
		res.Buttons[i] = &aimv1.Button{
			Id:        b.ID,
			Name:      b.Name,
			Tag:       b.Tag,
			HasIrdata: b.HasIRData,
		}
	}

	return
}
