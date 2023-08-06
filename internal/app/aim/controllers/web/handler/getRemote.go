package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
)

func (h *Handler) GetRemote(ctx context.Context, req *aimv1.GetRemoteRequest) (res *aimv1.GetRemoteResponse, err error) {
	var out bdy.GetRemoteOutput
	res = &aimv1.GetRemoteResponse{}
	in := bdy.GetRemoteInput{
		RemoteID: req.RemoteId,
	}

	out, err = h.i.GetRemote(ctx, in)
	if err != nil {
		return
	}

	res.Remote = &aimv1.Remote{
		Id:       out.Remote.ID,
		Name:     out.Remote.Name,
		Tag:      out.Remote.Tag,
		DeviceId: out.Remote.DeviceID,
	}
	return res, err
}
