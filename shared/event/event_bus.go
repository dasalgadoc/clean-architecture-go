package event

import "context"

type EventBus interface {
	Publish(context.Context, []Event) error
	Subscribe(Type, Handler)
}

type Handler interface {
	Handle(context.Context, Event) error
}
