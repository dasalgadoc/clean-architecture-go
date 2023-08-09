package domain

import "github.com/dasalgadoc/clean-architecture-go/shared/event"

const COURSE_CREATED_EVENT_NAME = "dasalgadoc.com.repository.1.course.created"

type CourseCreatedEvent struct {
	event.BaseEvent
	id   string
	name string
}

func (e CourseCreatedEvent) Type() event.Type {
	return COURSE_CREATED_EVENT_NAME
}

func (e CourseCreatedEvent) Id() string {
	return e.id
}

func (e CourseCreatedEvent) Name() string {
	return e.name
}

func NewCourseCreatedEvent(id, name string) CourseCreatedEvent {
	return CourseCreatedEvent{
		id:   id,
		name: name,

		BaseEvent: event.NewBaseEvent(id),
	}
}
