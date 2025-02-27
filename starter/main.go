package main

import (
	"await-signal"
	"await-signal/common"
	"context"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"log"
	"time"
)

func main() {
	c := common.GetEitherTemporalClient()
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "greetings_" + uuid.New().String(),
		TaskQueue: common.TaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, await_signal.Workflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)

	for err != nil {
		time.Sleep(2 * time.Second)
		//log.Println("Retry waiting for workflow completion")
		c.Close()
		c = common.GetEitherTemporalClient()
		err = c.GetWorkflow(context.Background(), we.GetID(), we.GetRunID()).Get(context.Background(), &result)
	}

	log.Println("Workflow result:", result)
}
