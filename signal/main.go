package main

import (
	await_signal "await-signal"
	"await-signal/common"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: program <workflow_id> <greeting_name>")
		return
	}

	workflowID := os.Args[1]
	greetingName := os.Args[2]

	c := common.GetEitherTemporalClient()
	defer c.Close()

	err := c.SignalWorkflow(context.Background(), workflowID, "", await_signal.SignalName, greetingName)
	if err != nil {
		log.Println("Unable to signal workflow.", err)
	} else {
		log.Println("Signaled workflow.")
	}
}
