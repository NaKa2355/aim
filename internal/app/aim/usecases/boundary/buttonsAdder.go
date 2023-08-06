package boundary

import "context"

type AddButtonInput struct {
	RemoteID string
	Tag      string
	Name     string
}

type AddButtonOutput struct {
	Button Button
}

type ButtonAdder interface {
	AddButton(ctx context.Context, i AddButtonInput) (AddButtonOutput, error)
}
