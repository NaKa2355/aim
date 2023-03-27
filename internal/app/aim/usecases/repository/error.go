package repository

type ErrCode string

const (
	CodeNotFound      = ErrCode("not found")
	CodeInvaildInput  = ErrCode("invaild input")
	CodeAlreadyExists = ErrCode("already exists")
	CodeDataBase      = ErrCode("database error")
)

type Error struct {
	Err  error
	Code ErrCode
}

func (e Error) Error() string {
	return e.Err.Error()
}

func NewError(code ErrCode, err error) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}
