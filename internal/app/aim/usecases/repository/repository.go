package repository

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/remote"
)

type Repository interface {
	CreateRemote(ctx context.Context, a *remote.Remote) (*remote.Remote, error)
	CreateButton(context.Context, remote.ID, *button.Button) (*button.Button, error)

	ReadRemote(context.Context, remote.ID) (*remote.Remote, error)
	ReadRemotes(context.Context) ([]*remote.Remote, error)
	ReadButtons(ctx context.Context, appID remote.ID) ([]*button.Button, error)
	ReadButton(context.Context, remote.ID, button.ID) (*button.Button, error)

	UpdateRemote(context.Context, *remote.Remote) error
	UpdateButton(context.Context, remote.ID, *button.Button) error

	DeleteRemote(context.Context, remote.ID) error
	DeleteButton(context.Context, remote.ID, button.ID) error
}
