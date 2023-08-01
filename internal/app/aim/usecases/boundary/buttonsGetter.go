package boundary

import "context"

type GetButtonsOutput struct {
	Buttons []Button
}

type GetButtonsInput struct {
	RemoteID string
}

type ButtonsGetter interface {
	GetButtons(ctx context.Context, i GetButtonsInput) (GetButtonsOutput, error)
}
