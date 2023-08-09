package event

import "context"

type EventBus interface {
	Publish(context.Context, []Event) error
}