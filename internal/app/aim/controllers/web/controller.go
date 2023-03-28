package web

import (
	"context"
	"sync"
)

type SessionID uint16

type ContextKey int

const SessionIDKey ContextKey = 0

type Response struct {
	Data any
	Err  error
}

type Controller struct {
	sessions map[SessionID](chan<- Response)
	mu       sync.RWMutex
	latestID SessionID
}

func NewController() *Controller {
	return &Controller{
		sessions: make(map[SessionID]chan<- Response),
		mu:       sync.RWMutex{},
		latestID: 0,
	}
}

func (c *Controller) NewSession(ctx context.Context) (context.Context, <-chan Response) {
	c.mu.Lock()
	defer c.mu.Unlock()
	id := c.latestID
	ch := make(chan Response)
	c.sessions[id] = ch
	ctx = context.WithValue(ctx, SessionIDKey, id)
	c.latestID++
	return ctx, ch
}

func (c *Controller) DeleteSession(id SessionID) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if s, ok := c.sessions[id]; ok {
		close(s)
		delete(c.sessions, id)
	}
}

func (c *Controller) SendResponse(id SessionID, resp Response) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if s, ok := c.sessions[id]; ok {
		s <- resp
	}
}
