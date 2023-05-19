package interactor

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

type Interactor struct {
	repo   repository.Repository
	output bdy.RemoteUpdateNotifier
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

func New(repo repository.Repository, out bdy.RemoteUpdateNotifier) *Interactor {
	i := &Interactor{
		repo:   repo,
		output: out,
	}
	return i
}

func (i *Interactor) AddRemote(ctx context.Context, in bdy.AddRemoteInput) (bdy.AddRemoteOutput, error) {
	out, err := i.addRemote(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) AddButton(ctx context.Context, in bdy.AddButtonInput) error {
	err := i.addButton(ctx, in)
	return wrapErr(err)
}

func (i *Interactor) GetRemotes(ctx context.Context) (bdy.GetRemotesOutput, error) {
	out, err := i.getRemotes(ctx)
	return out, wrapErr(err)
}

func (i *Interactor) GetRemote(ctx context.Context, in bdy.GetRemoteInput) (bdy.GetRemoteOutput, error) {
	out, err := i.getRemote(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) GetButtons(ctx context.Context, in bdy.GetButtonsInput) (bdy.GetButtonsOutput, error) {
	out, err := i.getButtons(ctx, in)
	return out, wrapErr(err)
}

func (i *Interactor) GetIRData(ctx context.Context, in bdy.GetIRDataInput) (bdy.GetIRDataOutput, error) {
	out, err := i.getIRData(ctx, in)
	return out, wrapErr(err)
}

// Update
func (i *Interactor) EditRemote(ctx context.Context, in bdy.EditRemoteInput) error {
	err := i.editRemote(ctx, in)
	return wrapErr(err)
}

func (i *Interactor) EditButton(ctx context.Context, in bdy.EditButtonInput) error {
	err := i.renameButton(ctx, in)
	return wrapErr(err)
}

func (i *Interactor) SetIRData(ctx context.Context, in bdy.SetIRDataInput) error {
	err := i.setIRData(ctx, in)
	return wrapErr(err)
}

// Delete
func (i *Interactor) DeleteRemote(ctx context.Context, in bdy.DeleteRemoteInput) error {
	err := i.deleteRemote(ctx, in)
	return wrapErr(err)
}

func (i *Interactor) DeleteButton(ctx context.Context, in bdy.DeleteButtonInput) error {
	err := i.deleteButton(ctx, in)
	return wrapErr(err)
}
