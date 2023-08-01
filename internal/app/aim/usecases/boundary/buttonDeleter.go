package boundary

import "context"

type DeleteButtonInput struct {
	RemoteID string
	ButtonID string
}

type ButtonDeleter interface {
	DeleteButton(ctx context.Context, i DeleteButtonInput) error
}
