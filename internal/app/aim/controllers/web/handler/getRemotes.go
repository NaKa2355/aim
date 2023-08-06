package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *Handler) GetRemotes(ctx context.Context, _ *empty.Empty) (res *aimv1.GetRemotesResponse, err error) {
	var out bdy.GetRemotesOutput
	res = &aimv1.GetRemotesResponse{}

	out, err = h.i.GetRemotes(ctx)
	if err != nil {
		return
	}

	res.Remotes = make([]*aimv1.Remote, len(out.Remotes))
	for i, r := range out.Remotes {
		res.Remotes[i] = &aimv1.Remote{
			Id:       r.ID,
			Name:     r.Name,
			Tag:      r.Tag,
			DeviceId: r.DeviceID,
		}
	}
	return
}
