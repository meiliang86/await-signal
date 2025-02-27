package main

import (
	"await-signal"
	"await-signal/common"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"log"
)

func main() {
	c := common.GetTemporalClient()
	defer c.Close()

	w := worker.New(c, common.TaskQueue, worker.Options{})

	w.RegisterWorkflowWithOptions(await_signal.Workflow, workflow.RegisterOptions{Name: "await_signal"})
	activities := &await_signal.Activities{Greeting: "Hello"}
	w.RegisterActivity(activities)

	err := w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
