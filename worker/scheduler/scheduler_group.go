package scheduler

import (
	"github.com/robfig/cron/v3"
)

// SchedulerGroup ...
type SchedulerGroup struct {
	id           cron.EntryID // the entry id of job
	spec         string       // the schedule string
	pl           ISchedulerTask
	retries      int    // the number of retries
	scheduleType string // the schedule type can be: once, endless
	finished     bool
	status       int
	name         string
}

// NewSchedulerGroup ...
func NewSchedulerGroup(cfg *ScheduleHandlerConfig) *SchedulerGroup {
	return &SchedulerGroup{
		spec:         cfg.Spec,
		pl:           cfg.Handler,
		retries:      cfg.Retries,
		scheduleType: cfg.Type,
		name:         cfg.Handler.GetName(),
	}
}
