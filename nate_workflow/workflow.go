package nate_workflow

import (
	"go.temporal.io/sdk/workflow"
	"natealcedo/temporal-workers/go/nate_activity"
	"time"
)

func Greet(ctx workflow.Context, greetMsg nate_activity.GreetMessage) (string, error) {
	workflow.Sleep(ctx, 10*time.Second) // Simulate a long-running operation
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, "SayHello", greetMsg).Get(ctx, &result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func FetchUsers(ctx workflow.Context) ([]nate_activity.User, error) {
	workflow.Sleep(ctx, 5*time.Second) // Simulate a long-running operation
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Minute, // Adjust the timeout as necessary
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var users []nate_activity.User
	err := workflow.ExecuteActivity(ctx, "FetchUsers").Get(ctx, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
