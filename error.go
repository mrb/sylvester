package sylvester

import (
	"fmt"
)

// Wrap error types with a struct that can hold the Node ID.
type EventError struct {
	Id  []byte // Node ID
	Err error  // Actual error
}

func NewEventError(id []byte, err error) *EventError {
	return &EventError{
		Id:  id,
		Err: err,
	}
}

func (e *EventError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Id, e.Err.Error())
}
