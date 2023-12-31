// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package todogen

import (
	"context"
	"emperror.dev/errors"
	"fmt"
	"github.com/piotr-rusin/playing-with-golang/internal/app/mga/todo"
)

// MarkedAsCompleteHandler handles MarkedAsComplete events.
type MarkedAsCompleteHandler interface {
	// MarkedAsComplete handles a(n) MarkedAsComplete event.
	MarkedAsComplete(ctx context.Context, event todo.MarkedAsComplete) error
}

// MarkedAsCompleteEventHandler handles MarkedAsComplete events.
type MarkedAsCompleteEventHandler struct {
	handler MarkedAsCompleteHandler
	name    string
}

// NewMarkedAsCompleteEventHandler returns a new MarkedAsCompleteEventHandler instance.
func NewMarkedAsCompleteEventHandler(handler MarkedAsCompleteHandler, name string) MarkedAsCompleteEventHandler {
	return MarkedAsCompleteEventHandler{
		handler: handler,
		name:    name,
	}
}

// HandlerName returns the name of the event handler.
func (h MarkedAsCompleteEventHandler) HandlerName() string {
	return h.name
}

// NewEvent returns a new empty event used for serialization.
func (h MarkedAsCompleteEventHandler) NewEvent() interface{} {
	return &todo.MarkedAsComplete{}
}

// Handle handles an event.
func (h MarkedAsCompleteEventHandler) Handle(ctx context.Context, event interface{}) error {
	e, ok := event.(*todo.MarkedAsComplete)
	if !ok {
		return errors.NewWithDetails("unexpected event type", "type", fmt.Sprintf("%T", event))
	}

	return h.handler.MarkedAsComplete(ctx, *e)
}
