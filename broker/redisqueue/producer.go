package redisqueue

import (
	"encoding/json"
	"fmt"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

// ISender ...
type ISender interface {
	SendBytes(msg []byte) error
	SendString(msg string) error
	SendDynamicQueue(dynamicTopic string, msg []byte) error
}

type IProducer interface {
	GetSender(topic string) ISender
}

// MQTTProducer ...
type redisQueueProducer struct {
	producer *work.Enqueuer
}

// NewProducer ...
func NewProducer(cfg *Config) (IProducer, error) {
	cfg.Prepare()

	// Make a redis pool
	var redisPool = &redis.Pool{
		MaxActive:   _defaultMaxActive,
		MaxIdle:     _defaultMaxIdle,
		IdleTimeout: _defaultIdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			options := []redis.DialOption{
				redis.DialDatabase(cfg.Database),
			}
			if cfg.Password != "" {
				options = append(options, redis.DialPassword(cfg.Password))
			}

			return redis.DialURL(cfg.Address, options...)
		},
	}

	// Make an enqueuer with a particular namespace
	var enqueuer = work.NewEnqueuer(cfg.QueuePrefix, redisPool)

	p := &redisQueueProducer{
		producer: enqueuer,
	}

	return p, nil
}

// GetSender ...
func (p *redisQueueProducer) GetSender(topic string) ISender {
	return &sender{p, topic}
}

type sender struct {
	*redisQueueProducer
	topic string
}

// SendBytes produce base-event into redis
func (p sender) SendBytes(data []byte) error {
	args := work.Q{}
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}

	// Enqueue a job named with the specified parameters.
	_, err := p.producer.Enqueue(p.topic, args)
	if err != nil {
		return err
	}

	return nil
}

// SendBytes produce base-event into dynamic queue redis
func (p sender) SendDynamicQueue(dynamicTopic string, data []byte) error {
	args := work.Q{}
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}
	dynamicQueue := fmt.Sprintf("%s:%s", p.topic, dynamicTopic)
	// Enqueue a dynamic job named with the specified parameters.
	_, err := p.producer.Enqueue(dynamicQueue, args)
	if err != nil {
		return err
	}

	return nil
}

// SendString produce string into mqtt
func (p sender) SendString(msg string) error {
	return p.SendBytes([]byte(msg))
}
