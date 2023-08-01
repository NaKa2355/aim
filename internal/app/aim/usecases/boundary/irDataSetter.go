package boundary

import "context"

type SetIRDataInput struct {
	RemoteID string
	ButtonID string
	Data     IRData
}

type IRDataSetter interface {
	SetIRData(ctx context.Context, i SetIRDataInput) error
}
