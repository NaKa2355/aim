package boundary

import "context"

type GetIRDataOutput struct {
	IRData IRData
}

type GetIRDataInput struct {
	RemoteID string
	ButtonID string
}

type IRDataGetter interface {
	GetIRData(ctx context.Context, i GetIRDataInput) (GetIRDataOutput, error)
}
