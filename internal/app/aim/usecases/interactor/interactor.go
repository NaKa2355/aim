package interactor

import (
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

type Interactor struct {
	repo   repository.Repository
	output bdy.RemoteUpdateNotifier
}

func New(repo repository.Repository, out bdy.RemoteUpdateNotifier) *Interactor {
	i := &Interactor{
		repo:   repo,
		output: out,
	}
	return i
}
