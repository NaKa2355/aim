package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func (h *Handler) GetIrData(ctx context.Context, req *aimv1.GetIrDataRequest) (res *anypb.Any, err error) {
	var in bdy.GetIRDataInput
	var out bdy.GetIRDataOutput
	res = &anypb.Any{}

	in.RemoteID = req.RemoteId
	in.ButtonID = req.ButtonId
	out, err = h.i.GetIRData(ctx, in)
	if err != nil {
		return
	}

	proto.Unmarshal(out.IRData, res)
	return
}
