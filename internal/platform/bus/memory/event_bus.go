package memory

import (
	"context"
	"github.com/dasalgadoc/clean-architecture-go/shared/event"
)

type EventBus struct {
	events []event.Event
}

func NewEventBus() *EventBus {
	return &EventBus{}
}

func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	b.events = append(b.events, events...)
	return nil
}
