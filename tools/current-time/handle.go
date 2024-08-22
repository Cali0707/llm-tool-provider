package function

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
)

type CurrentTime struct {
	Now string `json:"now"`
}

// Handle an event.
func Handle(ctx context.Context, e event.Event) (*event.Event, error) {
	fmt.Printf("Received a new event\n%s", e.String())

	response := event.New()
	response.SetID(uuid.New().String())
	response.SetType("current.time.response")
	response.SetSource("/current-time")
	response.SetExtension("responseid", e.ID())
	err := response.SetData(event.TextJSON, CurrentTime{Now: time.Now().Format(time.RFC3339)})
	if err != nil {
		fmt.Printf("failed to set data on event: %s", err.Error())
		return nil, err
	}
	return &response, nil // echo to caller
}

/*
Other supported function signatures:

	Handle()
	Handle() error
	Handle(context.Context)
	Handle(context.Context) error
	Handle(event.Event)
	Handle(event.Event) error
	Handle(context.Context, event.Event)
	Handle(context.Context, event.Event) error
	Handle(event.Event) *event.Event
	Handle(event.Event) (*event.Event, error)
	Handle(context.Context, event.Event) *event.Event
	Handle(context.Context, event.Event) (*event.Event, error)

*/
