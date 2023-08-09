package memory

import (
	"context"
	"github.com/dasalgadoc/clean-architecture-go/shared/event"
)

type EventBus struct {
	handlers map[event.Type][]event.Handler
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event.Type][]event.Handler),
	}
}

func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers[evt.Type()]
		if !ok {
			return nil
		}

		for _, handler := range handlers {
			err := handler.Handle(ctx, evt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *EventBus) Subscribe(eventType event.Type, handler event.Handler) {
	subscriberForType, ok := b.handlers[eventType]
	if !ok {
		b.handlers[eventType] = []event.Handler{handler}
	}
	subscriberForType = append(subscriberForType, handler)
}
