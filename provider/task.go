package provider

import (
	"context"
	"fmt"
	"github.com/dienggo/diego/app/background"
	"github.com/dienggo/diego/pkg/task"
	"os"
	"os/signal"
	"syscall"
)

// Don't remove or edit manually all comments on this method
// tasks : Slice of task registry
func tasks() []task.ITask {
	return []task.ITask{ // List-Of-Generated-Task-By-Diego
		background.Kernel{},
	}
}

func NewAppTaskProvider() IProvider {
	return &appTask{}
}

type appTask struct{}

func (t appTask) Provide() {
	go runTask()
}

// runTask : run task on other thread
func runTask() {
	ctx := context.Background()
	for _, iTask := range tasks() {
		go task.Runner(iTask, ctx, false)
	}
	// Listen for shutdown signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx.Done()
	fmt.Println("All tasks shutting down completed")
}
