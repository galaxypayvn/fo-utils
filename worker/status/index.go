package workerstatus

// ProcessStatus ...
type ProcessStatus struct {
	Code    Status
	Message []byte
}

var (
	ProcessOK            = ProcessStatus{Code: Success}
	ProcessFailRetry     = ProcessStatus{Code: Retry}
	ProcessFailDrop      = ProcessStatus{Code: Drop}
	ProcessFailReproduce = ProcessStatus{Code: FailReproduce}
)
