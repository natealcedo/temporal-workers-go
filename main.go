package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"natealcedo/temporal-workers/go/nate_activity"
	"natealcedo/temporal-workers/go/nate_workflow"
)

func main() {
	// Create a new client to connect to your Temporal server
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	// Create a new worker that listens on the "greet-tq" task queue
	w := worker.New(c, "greet-tq", worker.Options{})

	// Register your workflow and activity with the worker
	w.RegisterWorkflow(nate_workflow.Greet)
	w.RegisterWorkflow(nate_workflow.FetchUsers)
	w.RegisterActivity(&nate_activity.GreetActivitiesImpl{})

	// Register your workflow and activity with the worker
	// Start the worker in a separate goroutine so it doesn't block
	go func() {
		if err := w.Run(worker.InterruptCh()); err != nil {
			log.Fatalln("Unable to start worker", err)
		}
	}()

	// Prevent the main function from returning, which would stop the worker
	select {} // Block forever
}
