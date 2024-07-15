package nate_activity

import (
	"context"
)

type GreetMessage struct {
	Name string `json:"name"`
}

type GreetActivitiesImpl struct{}

func (a *GreetActivitiesImpl) SayHello(ctx context.Context, greetMsg GreetMessage) (string, error) {

	// Use the unmarshaled data
	return "Hello, " + greetMsg.Name, nil
}
