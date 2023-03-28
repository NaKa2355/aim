package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

type Interactor struct {
	repo repository.Repository
}

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

func New(in repository.Repository) *Interactor {
	i := &Interactor{
		repo: in,
	}
	return i
}

func (i *Interactor) AddCustom(ctx context.Context, in bdy.AddCustomInput) (bdy.AddAppOutput, error) {
	out, err := i.addCustom(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) AddToggle(ctx context.Context, in bdy.AddToggleInput) (bdy.AddAppOutput, error) {
	out, err := i.addToggle(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) AddButton(ctx context.Context, in bdy.AddButtonInput) (bdy.AddAppOutput, error) {
	out, err := i.addButton(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) AddThermostat(ctx context.Context, in bdy.AddThermostatInput) (bdy.AddAppOutput, error) {
	out, err := i.addThermostat(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) AddCommand(ctx context.Context, in bdy.AddCommandInput) error {
	err := i.addCommand(ctx, in)
	return wrapErr(err)
}

// Read
func (i *Interactor) GetCustom(ctx context.Context, in bdy.GetAppInput) (bdy.GetCustomOutput, error) {
	out, err := i.getCustom(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) GetToggle(ctx context.Context, in bdy.GetAppInput) (bdy.GetToggleOutput, error) {
	out, err := i.getToggle(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) GetButton(ctx context.Context, in bdy.GetAppInput) (bdy.GetButtonOutput, error) {
	out, err := i.getButton(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) GetThermostat(ctx context.Context, in bdy.GetAppInput) (bdy.GetThermostatOutput, error) {
	out, err := i.getThermostat(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) GetAppliances(ctx context.Context) (bdy.GetAppliancesOutput, error) {
	out, err := i.getAppliances(ctx)
	return out, wrapErr(err)
}

func (i *Interactor) GetCommand(ctx context.Context, in bdy.GetCommandInput) (bdy.GetCommandOutput, error) {
	out, err := i.getCommand(ctx, in)
	return out, wrapErr(err)
}

// Update
func (i *Interactor) RenameAppliance(ctx context.Context, in bdy.RenameAppInput) error {
	err := i.renameAppliance(ctx, in)
	return wrapErr(err)
}

func (i *Interactor) ChangeIRDevice(ctx context.Context, in bdy.ChangeIRDevInput) error {
	err := i.changeIRDevice(ctx, in)
	return wrapErr(err)
}

func (i *Interactor) RenameCommand(ctx context.Context, in bdy.RenameCommandInput) error {
	err := i.renameCommand(ctx, in)
	return wrapErr(err)
}

func (i *Interactor) SetIRData(ctx context.Context, in bdy.SetIRDataInput) error {
	err := i.setIRData(ctx, in)
	return wrapErr(err)
}

// Delete
func (i *Interactor) DeleteAppliance(ctx context.Context, in bdy.DeleteAppInput) error {
	err := i.deleteAppliance(ctx, in)
	return wrapErr(err)
}

func (i *Interactor) DeleteCommand(ctx context.Context, in bdy.DeleteCommandInput) error {
	err := i.deleteCommand(ctx, in)
	return wrapErr(err)
}
