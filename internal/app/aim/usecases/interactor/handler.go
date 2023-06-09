package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

type Interactor struct {
	repo   repository.Repository
	output bdy.ApplianceUpdateNotifier
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

func New(repo repository.Repository, out bdy.ApplianceUpdateNotifier) *Interactor {
	i := &Interactor{
		repo:   repo,
		output: out,
	}
	return i
}

func (i *Interactor) AddAppliance(ctx context.Context, in bdy.AddApplianceInput) (bdy.AddAppOutput, error) {
	out, err := i.addAppliance(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) AddCommand(ctx context.Context, in bdy.AddCommandInput) error {
	err := i.addCommand(ctx, in)
	return wrapErr(err)
}

func (i *Interactor) GetAppliances(ctx context.Context) (bdy.GetAppliancesOutput, error) {
	out, err := i.getAppliances(ctx)
	return out, wrapErr(err)
}

func (i *Interactor) GetAppliance(ctx context.Context, in bdy.GetApplianceInput) (bdy.GetApplianceOutput, error) {
	out, err := i.getAppliance(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) GetCommands(ctx context.Context, in bdy.GetCommandsInput) (bdy.GetCommandsOutput, error) {
	out, err := i.getCommands(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) GetIRData(ctx context.Context, in bdy.GetIRDataInput) (bdy.GetIRDataOutput, error) {
	out, err := i.getIRData(ctx, in)
	return out, wrapErr(err)
}

// Update
func (i *Interactor) EditAppliance(ctx context.Context, in bdy.EditApplianceInput) error {
	err := i.editAppliance(ctx, in)
	return wrapErr(err)
}

func (i *Interactor) EditCommand(ctx context.Context, in bdy.EditCommandInput) error {
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
