package boundary

import "context"

type AddButtonInput struct {
	RemoteID string
	Name     string
}
type ButtonAdder interface {
	AddButton(ctx context.Context, i AddButtonInput) error
}
