package boundary

import "context"

type EditButtonInput struct {
	RemoteID string
	ButtonID string
	Name     string
}

type ButtonEditor interface {
	EditButton(ctx context.Context, i EditButtonInput) error
}
