package boundary

type ErrCode string

const (
	CodeNotFound         = ErrCode("not found")
	CodeInvaildInput     = ErrCode("invaild input")
	CodeAlreadyExists    = ErrCode("already exists")
	CodeInvaildOperation = ErrCode("invaild operation")
	CodeDatabase         = ErrCode("database error")
)

type Error struct {
	Code ErrCode
	Err  error
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
