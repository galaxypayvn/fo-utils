package workerstatus

// ProcessStatus ...
type ProcessStatus struct {
	Code    Status
	Message string
}

var (
	ProcessOK            = ProcessStatus{Code: Success}
	ProcessFailRetry     = ProcessStatus{Code: Retry}
	ProcessFailDrop      = ProcessStatus{Code: Drop}
	ProcessFailReproduce = ProcessStatus{Code: FailReproduce}
)

// WithMessage adds a custom message to ProcessOK status
func (ps ProcessStatus) WithMessage(message string) ProcessStatus {
	return ProcessStatus{
		Code:    ps.Code,
		Message: message,
	}
}
