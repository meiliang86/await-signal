package await_signal

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

const SignalName = "greeting_name"

func Workflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var a *Activities // use a nil struct pointer to call activities that are part of a structure

	var greetResult string
	err := workflow.ExecuteActivity(ctx, a.GetGreeting).Get(ctx, &greetResult)
	if err != nil {
		logger.Error("Get greeting failed.", "Error", err)
		return "", err
	}

	var greetingName string
	workflow.GetSignalChannel(ctx, SignalName).Receive(ctx, &greetingName)

	// Say Greeting.
	var sayResult string
	err = workflow.ExecuteActivity(ctx, a.SayGreeting, greetResult, greetingName).Get(ctx, &sayResult)
	if err != nil {
		logger.Error("Marshalling failed with error.", "Error", err)
		return "", err
	}

	return sayResult, nil
}
