package task

import "context"

type Scheduler struct {
	overlaps bool
	isBusy   bool
	duration *Duration
	name     string
	Run      func(ctx context.Context)
}

// Name is method to set name of scheduler
func (s Scheduler) Name(name string) *Scheduler {
	s.name = name
	return &s
}

// Every is method to set task run delayed
func (s Scheduler) Every(duration Duration) *Scheduler {
	s.duration = &duration
	return &s
}

// Overlaps is method to set task run with overlaps or not
func (s Scheduler) Overlaps(overlaps bool) *Scheduler {
	s.overlaps = overlaps
	return &s
}
