package event

import (
	"github.com/dasalgadoc/clean-architecture-go/shared/domain"
	"log"
	"time"
)

type Type string

type Event interface {
	ID() string
	AggregateID() string
	OccurredAt() time.Time
	Type() Type
}

type BaseEvent struct {
	eventId     domain.UUIDValueObject
	aggregateId string
	occurredAt  time.Time
}

func (b BaseEvent) ID() string {
	return b.eventId.Value()
}

func (b BaseEvent) AggregateID() string {
	return b.aggregateId
}

func (b BaseEvent) OccurredAt() time.Time {
	return b.occurredAt
}

func NewBaseEvent(aggregateId string) BaseEvent {
	uid, err := domain.NewUUIDValueObject()
	if err != nil {
		log.Fatal(err)
	}
	return BaseEvent{
		eventId:     *uid,
		aggregateId: aggregateId,
		occurredAt:  time.Now(),
	}
}
