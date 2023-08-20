package handler

import (
	"context"
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
	i                Boundary
	nc               chan aimv1.UpdateNotification
	notification     aimv1.UpdateNotification
	c                *cond.Cond
	StreamingContext context.Context
}

var _ aimv1.AimServiceServer = &Handler{}

func New(ctx context.Context) *Handler {
	return &Handler{
		nc:               make(chan aimv1.UpdateNotification),
		c:                cond.NewCond(&sync.Mutex{}),
		StreamingContext: ctx,
	}
}

func (h *Handler) SetInteractor(i Boundary) {
	h.i = i
}
