package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

type Interactor struct {
	repo   repository.Repository
	output bdy.OutputBoundary
}

var _ bdy.InputBoundary = &Interactor{}

func convertErrCode(err error) bdy.ErrCode {
	var code bdy.ErrCode

	if entityErr, ok := err.(entities.Error); ok {
		switch entityErr.Code {
		case entities.CodeInvaildInput:
			code = bdy.CodeInvaildInput
		case entities.CodeInvaildOperation:
			code = bdy.CodeInvaildOperation
		}
	}

	if repoErr, ok := err.(repository.Error); ok {
		switch repoErr.Code {
		case repository.CodeInvaildInput:
			code = bdy.CodeInvaildInput
		case repository.CodeNotFound:
			code = bdy.CodeNotFound
		case repository.CodeDataBase:
			code = bdy.CodeDatabase
		case repository.CodeAlreadyExists:
			code = bdy.CodeAlreadyExists
		}
	}
	return code
}

func wrapErr(err error) error {
	if err == nil {
		return nil
	}

	code := convertErrCode(err)
	return bdy.NewError(code, err)
}

func New(in repository.Repository, o bdy.OutputBoundary) *Interactor {
	i := &Interactor{
		repo:   in,
		output: o,
	}
	return i
}

func (i *Interactor) AddCustom(ctx context.Context, in bdy.AddCustomInput) {
	out, err := i.addCustom(ctx, in)
	i.output.AddCustom(ctx, out, wrapErr(err))
	if err == nil {
		i.output.ChangeNotify(bdy.ChangeNotifyOutput{
			AppID: string(out.ID),
			Type:  bdy.ChangeTypeAdd,
		})
	}
}

func (i *Interactor) AddToggle(ctx context.Context, in bdy.AddToggleInput) {
	out, err := i.addToggle(ctx, in)
	i.output.AddToggle(ctx, out, wrapErr(err))
	if err == nil {
		i.output.ChangeNotify(bdy.ChangeNotifyOutput{
			AppID: string(out.ID),
			Type:  bdy.ChangeTypeAdd,
		})
	}
}

func (i *Interactor) AddButton(ctx context.Context, in bdy.AddButtonInput) {
	out, err := i.addButton(ctx, in)
	i.output.AddButton(ctx, out, wrapErr(err))
	if err == nil {
		i.output.ChangeNotify(bdy.ChangeNotifyOutput{
			AppID: string(out.ID),
			Type:  bdy.ChangeTypeAdd,
		})
	}
}

func (i *Interactor) AddThermostat(ctx context.Context, in bdy.AddThermostatInput) {
	out, err := i.addThermostat(ctx, in)
	i.output.AddThermostat(ctx, out, wrapErr(err))
	if err == nil {
		i.output.ChangeNotify(bdy.ChangeNotifyOutput{
			AppID: string(out.ID),
			Type:  bdy.ChangeTypeAdd,
		})
	}
}

func (i *Interactor) AddCommand(ctx context.Context, in bdy.AddCommandInput) {
	err := i.addCommand(ctx, in)
	i.output.AddCommand(ctx, wrapErr(err))
}

// Read
func (i *Interactor) GetCustom(ctx context.Context, in bdy.GetAppInput) {
	out, err := i.getCustom(ctx, in)
	i.output.GetCustom(ctx, out, wrapErr(err))
}

func (i *Interactor) GetToggle(ctx context.Context, in bdy.GetAppInput) {
	out, err := i.getToggle(ctx, in)
	i.output.GetToggle(ctx, out, wrapErr(err))
}

func (i *Interactor) GetButton(ctx context.Context, in bdy.GetAppInput) {
	out, err := i.getButton(ctx, in)
	i.output.GetButton(ctx, out, wrapErr(err))
}

func (i *Interactor) GetThermostat(ctx context.Context, in bdy.GetAppInput) {
	out, err := i.getThermostat(ctx, in)
	i.output.GetThermostat(ctx, out, wrapErr(err))
}

func (i *Interactor) GetAppliances(ctx context.Context) {
	out, err := i.getAppliances(ctx)
	i.output.GetAppliances(ctx, out, wrapErr(err))
}

func (i *Interactor) GetCommand(ctx context.Context, in bdy.GetCommandInput) {
	out, err := i.getCommand(ctx, in)
	i.output.GetCommand(ctx, out, wrapErr(err))
}

// Update
func (i *Interactor) RenameAppliance(ctx context.Context, in bdy.RenameAppInput) {
	err := i.renameAppliance(ctx, in)
	i.output.RenameAppliance(ctx, wrapErr(err))
}

func (i *Interactor) ChangeIRDevice(ctx context.Context, in bdy.ChangeIRDevInput) {
	err := i.changeIRDevice(ctx, in)
	i.output.ChangeIRDevice(ctx, wrapErr(err))
}

func (i *Interactor) RenameCommand(ctx context.Context, in bdy.RenameCommandInput) {
	err := i.renameCommand(ctx, in)
	i.output.RenameCommand(ctx, wrapErr(err))
}

func (i *Interactor) SetIRData(ctx context.Context, in bdy.SetIRDataInput) {
	err := i.setIRData(ctx, in)
	i.output.SetIRData(ctx, wrapErr(err))
}

// Delete
func (i *Interactor) DeleteAppliance(ctx context.Context, in bdy.DeleteAppInput) {
	err := i.deleteAppliance(ctx, in)
	i.output.DeleteAppliance(ctx, wrapErr(err))
}

func (i *Interactor) DeleteCommand(ctx context.Context, in bdy.DeleteCommandInput) {
	err := i.deleteCommand(ctx, in)
	i.output.DeleteCommand(ctx, wrapErr(err))
}
