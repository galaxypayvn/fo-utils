package messagecode

const (
	GeneralSuccessCreatedCode = 102001
	GeneralSuccessGetCode     = 102000
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

func (err Error) Error() string {
	return err.Cause.Error()
}
