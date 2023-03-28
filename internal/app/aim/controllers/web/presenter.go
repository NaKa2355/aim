package web

import (
	"context"

	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
)

func (c *Controller) AddCustom(ctx context.Context, out bdy.AddAppOutput, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Data: out, Err: err})
	c.DeleteSession(id)
}

func (c *Controller) AddToggle(ctx context.Context, out bdy.AddAppOutput, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Data: out, Err: err})
	c.DeleteSession(id)
}

func (c *Controller) AddButton(ctx context.Context, out bdy.AddAppOutput, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Data: out, Err: err})
	c.DeleteSession(id)
}

func (c *Controller) AddThermostat(ctx context.Context, out bdy.AddAppOutput, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Data: out, Err: err})
	c.DeleteSession(id)
}

func (c *Controller) AddCommand(ctx context.Context, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Err: err})
	c.DeleteSession(id)
}

// Read
func (c *Controller) GetCustom(ctx context.Context, out bdy.GetCustomOutput, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Err: err})
	c.DeleteSession(id)
}

func (c *Controller) GetToggle(ctx context.Context, out bdy.GetToggleOutput, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Data: out, Err: err})
	c.DeleteSession(id)
}

func (c *Controller) GetButton(ctx context.Context, out bdy.GetButtonOutput, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Data: out, Err: err})
	c.DeleteSession(id)
}

func (c *Controller) GetThermostat(ctx context.Context, out bdy.GetThermostatOutput, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Data: out, Err: err})
	c.DeleteSession(id)
}

func (c *Controller) GetAppliances(ctx context.Context, out bdy.GetAppliancesOutput, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Data: out, Err: err})
	c.DeleteSession(id)
}

func (c *Controller) GetCommand(ctx context.Context, out bdy.GetCommandOutput, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Data: out, Err: err})
	c.DeleteSession(id)
}

// Update
func (c *Controller) RenameAppliance(ctx context.Context, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Err: err})
	c.DeleteSession(id)
}

func (c *Controller) ChangeIRDevice(ctx context.Context, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Err: err})
	c.DeleteSession(id)
}

func (c *Controller) RenameCommand(ctx context.Context, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Err: err})
	c.DeleteSession(id)
}

func (c *Controller) SetIRData(ctx context.Context, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Err: err})
	c.DeleteSession(id)
}

// Delete
func (c *Controller) DeleteAppliance(ctx context.Context, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Err: err})
	c.DeleteSession(id)
}

func (c *Controller) DeleteCommand(ctx context.Context, err error) {
	id := ctx.Value(SessionIDKey).(SessionID)
	c.SendResponse(id, Response{Err: err})
	c.DeleteSession(id)
}

func (c *Controller) ChangeNotify(out bdy.ChangeNotifyOutput) {

}
