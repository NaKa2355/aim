package interactor

import (
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

type Interactor struct {
	rep repository.Repository
}

func New(r repository.Repository) *Interactor {
	i := &Interactor{
		rep: r,
	}
	return i
}
