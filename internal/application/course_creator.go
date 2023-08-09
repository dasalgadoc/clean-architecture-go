package application

import (
	"context"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"github.com/dasalgadoc/clean-architecture-go/shared/event"
)

type CourseCreator struct {
	courseRepository domain.CourseRepository
	eventBus         event.EventBus
}

func (cc *CourseCreator) Invoke(ctx context.Context, name string) error {
	course, err := domain.CreateCourse(name)
	if err != nil {
		return err
	}

	err = cc.courseRepository.Save(ctx, *course)
	if err != nil {
		return err
	}

	return cc.eventBus.Publish(ctx, course.PullEvents())
}

func NewCourseCreator(r domain.CourseRepository, bus event.EventBus) CourseCreator {
	return CourseCreator{
		courseRepository: r,
		eventBus:         bus,
	}
}
