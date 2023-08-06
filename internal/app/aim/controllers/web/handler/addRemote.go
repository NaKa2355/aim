package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
)

func (h *Handler) AddRemote(ctx context.Context, _req *aimv1.AddRemoteRequest) (res *aimv1.AddRemoteResponse, err error) {
	var in bdy.AddRemoteInput
	var out bdy.AddRemoteOutput

	in = bdy.AddRemoteInput{
		Name:     _req.Name,
		DeviceID: _req.DeviceId,
		Tag:      _req.Tag,
	}

	in.Buttons = make([]bdy.AddButtonInput, len(_req.Buttons))

	for i, b := range _req.Buttons {
		in.Buttons[i].Name = b.Name
		in.Buttons[i].Tag = b.Tag
	}

	out, err = h.i.AddRemote(ctx, in)
	if err != nil {
		return
	}

	res = &aimv1.AddRemoteResponse{
		RemoteId: out.Remote.ID,
	}
	return
}
