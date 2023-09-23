package background

import (
	"context"
	"github.com/dienggo/diego/pkg/task"
	"time"
)

type Kernel struct{}

// Task is main method to handle when task executed
func (t Kernel) Task() task.Scheduler {
	// sleep delay run every 10 second after method completed / without overlapping
	// use .Overlaps(true) to set with overlapping
	duration := 10 * time.Second
	return *task.Scheduler{Run: func(ctx context.Context) {
		// TODO : write your logic here

	}}.Every(task.Duration{SleepDuration: &duration}).Overlaps(false).Name("DefaultKernel")
}
