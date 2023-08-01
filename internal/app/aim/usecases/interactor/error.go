package interactor

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities"
	bdy "github.com/NaKa2355/aim/internal/app/aim/usecases/boundary"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

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
