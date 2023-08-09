package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"github.com/dasalgadoc/clean-architecture-go/shared/event"
)

type CourseCounterEventHandler struct {
}

func NewCourseCounter() CourseCounterEventHandler {
	return CourseCounterEventHandler{}
}

func (c CourseCounterEventHandler) Handle(_ context.Context, event event.Event) error {
	courseCreatedEvt, ok := event.(domain.CourseCreatedEvent)
	if !ok {
		return errors.New("invalid event type")
	}
	fmt.Println("Course created: ", courseCreatedEvt.Id(), courseCreatedEvt.Name())
	return nil
}
