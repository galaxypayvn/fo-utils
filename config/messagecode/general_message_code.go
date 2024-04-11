package messagecode

const (
	GeneralSuccessCreatedCode = 102001
	GeneralSuccessCode        = 102000
	GeneralInternalErrorCode  = 105000
	GeneralBadRequestCode     = 104000
	GeneralNotFoundCode       = 104004
	GeneralDuplicateRequest   = 104100
)

type Error struct {
	Code   int
	Params []any
	Cause  error
}

func NewError(code int, cause error, params ...any) Error {
	return Error{
		Code:   code,
		Params: params,
		Cause:  cause,
	}
}

func (err Error) Error() string {
	return err.Cause.Error()
}

func (err Error) Unwrap() error {
	return err.Cause
}
