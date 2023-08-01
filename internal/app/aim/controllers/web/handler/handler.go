package handler

import (
	"sync"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/pkg/cond"
	aimv1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
)

type Boundary interface {
	bdy.RemoteAdder
	bdy.ButtonAdder

	bdy.RemotesGetter
	bdy.RemoteGetter
	bdy.IRDataGetter
	bdy.ButtonsGetter

	bdy.RemoteEditor
	bdy.ButtonEditor
	bdy.IRDataSetter

	bdy.RemoteDeleter
	bdy.ButtonDeleter
}

type Handler struct {
	aimv1.UnimplementedAimServiceServer
	i            Boundary
	nc           chan aimv1.RemoteUpdateNotification
	notification aimv1.RemoteUpdateNotification
	c            *cond.Cond
}

var _ aimv1.AimServiceServer = &Handler{}

func New() *Handler {
	return &Handler{
		nc: make(chan aimv1.RemoteUpdateNotification),
		c:  cond.NewCond(&sync.Mutex{}),
	}
}

func (h *Handler) SetInteractor(i Boundary) {
	h.i = i
}
