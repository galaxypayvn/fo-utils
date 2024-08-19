package redisqueue

import (
	"fmt"
	"github.com/gocraft/work"
	"strings"
	"time"
)

const (
	_defaultConcurrency = 100 // tổng số công việc có thể được xử lý đồng thời bởi pool
	_defaultMaxActive   = 20
	_defaultMaxIdle     = 20
	_defaultIdleTimeout = 240 * time.Second
)

type JobOptions struct {
	Priority       uint // Priority from 1 to 10000 (biến này chỉ có hiệu lực khi 1 pool có nhiều loại job)
	MaxConcurrency uint // số worker tối đa hoạt động đồng thời xử lý 1 loại job (default is 0, meaning no max) (biến này chỉ có hiệu lực khi 1 pool có nhiều loại job)
	MaxFails       uint // cho phép một job tối đa mấy lần lỗi, biến này chỉ có hiệu lực khi handler trả về trạng thái muốn retry ( nếu set 1 thi gửi thẳng tới queue dead ko retry (unless SkipDead)
	SkipDead       bool // If true, don't send failed jobs to the dead queue when retries are exhausted.
}

type JobExecution struct {
	JobID      string `json:"job_id"`
	JobName    string `json:"job_name"`
	Args       work.Q `json:"args"`
	EnqueuedAt int64  `json:"enqueued_at"` // Thời gian (dưới dạng timestamp) khi công việc được thêm vào hàng đợi. Thời gian này giúp xác định khi nào công việc đã được đưa vào hệ thống.
	Unique     bool   `json:"unique"`      // Một cờ để xác định xem công việc có phải là duy nhất không. Điều này có thể được sử dụng để đảm bảo rằng không có công việc nào giống nhau được thêm vào hàng đợi.

	// Inputs when retrying
	Fails     int64  `json:"fails"` // number of times this job has failed
	FailedAt  int64  `json:"failed_at"`
	Error     string `json:"error"`
	LastError string `json:"last_error"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
}

type Config struct {
	Address         string   `json:"address" mapstructure:"address"`
	Password        string   `json:"password" mapstructure:"password"`
	Database        int      `json:"database" mapstructure:"database"`
	QueuePrefix     string   `json:"queue_prefix" mapstructure:"queue_prefix"`           // tên prefix cho queue
	QueueLog        string   `json:"queue_log" mapstructure:"queue_log"`                 // nếu dc set các job sẽ đc chạy qua 1 middle log tiến đô xử lý của 1 job
	QueueLogExclude []string `json:"queue_log_exclude" mapstructure:"queue_log_exclude"` // nếu dc set các job này sẽ ko chạy qua middle log
}

func (o *Config) Prepare() {
	if o == nil {
		return
	}

	if !strings.HasPrefix(o.Address, "redis://") {
		o.Address = fmt.Sprintf("redis://%v", o.Address)
	}

	if o.QueuePrefix == "" {
		o.QueuePrefix = "queue"
	}

}
