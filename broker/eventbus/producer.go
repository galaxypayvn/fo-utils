package eventbus

import (
	"github.com/asaskevich/EventBus"
)

// IProducer
type IProducer interface {
	GetSender(topic string) ISender
}

// ISender ...
type ISender interface {
	Emit(data interface{}) error
}

type messageInfo struct {
	msg        []byte
	routingKey string
	xchName    string
	mandatory  bool
	immediate  bool
	retries    int64
}

// eventBusProducer ...
type eventBusProducer struct {
	producer EventBus.Bus
}

// NewEventBusProducer ...
func NewEventBusProducer() IProducer {
	pObj := &eventBusProducer{
		producer: EventBus.New(),
	}

	return pObj
}

// GetSender ...
func (p *eventBusProducer) GetSender(topic string) ISender {
	return &sender{p, topic}
}

type sender struct {
	*eventBusProducer
	topic string
}

func (s sender) Emit(data interface{}) error {
	s.producer.Publish(s.topic, data)
	return nil
}
