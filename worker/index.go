package worker

import (
	"code.finan.cc/finan-one-be/fo-utils/worker/consumer"
	workerstatus "code.finan.cc/finan-one-be/fo-utils/worker/status"
)

// HandlerWOption ...
type HandlerWOption struct {
	MsgHandler
	Replica int64
}

// MsgHandler ...
type MsgHandler interface {
	Handle(msg []byte) workerstatus.ProcessStatus
}
type IWorker interface {
	Start()
}

// Worker ...
type worker struct {
	consumer *consumer.Consumer
}

func (w *worker) Start() {
	if w.consumer != nil {
		go w.consumer.Start()
	}
}
