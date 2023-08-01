package handler

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/proto"
)

func (h *Handler) SetIrData(ctx context.Context, req *aimv1.SetIRDataRequest) (e *empty.Empty, err error) {
	var in bdy.SetIRDataInput
	e = &empty.Empty{}

	in.RemoteID = req.RemoteId
	in.ButtonID = req.ButtonId
	in.Data, err = proto.Marshal(req.Irdata)
	if err != nil {
		return
	}

	err = h.i.SetIRData(ctx, in)
	return
}
