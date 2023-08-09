package mocks

import (
	"context"
	"github.com/dasalgadoc/clean-architecture-go/shared/event"
	"github.com/stretchr/testify/mock"
)

type EventBusMock struct {
	mock.Mock
}

func (e *EventBusMock) Publish(_ context.Context, event []event.Event) error {
	args := e.Called()
	return args.Error(0)
}
