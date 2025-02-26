package main

import (
	await_signal "await-signal"
	"await-signal/common"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

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
