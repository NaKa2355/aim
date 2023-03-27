package entities

type ErrCode string

const (
	CodeInvaildInput     = ErrCode("invaild input")
	CodeInvaildOperation = ErrCode("invaild operation")
)

type Error struct {
	Err  error
	Code ErrCode
}

func NewError(code ErrCode, err error) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func (e Error) Error() string {
	return e.Err.Error()
}
