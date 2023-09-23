package provider

import (
	"context"
	"github.com/dienggo/diego/app/background"
	"github.com/dienggo/diego/pkg/task"
)

// Don't remove or edit manually all comments on this method
// tasks : Slice of task registry
func tasks() []task.ITask {
	return []task.ITask{ // List-Of-Generated-Task-By-Diego
		background.Kernel{},
	}
}

type appTask struct{}

func (t appTask) Provide() {
	for _, iTask := range tasks() {
		go func() { task.Runner(iTask, context.Background(), false) }()
	}
}
