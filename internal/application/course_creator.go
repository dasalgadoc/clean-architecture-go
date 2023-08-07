package application

import (
	"context"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
)

type CourseCreator struct {
	courseRepository domain.CourseRepository
}

func (cc *CourseCreator) Invoke(ctx context.Context, name string) error {
	course, err := domain.CreateCourse(name)
	if err != nil {
		return err
	}

	return cc.courseRepository.Save(ctx, *course)
}

func NewCourseCreator(r domain.CourseRepository) CourseCreator {
	return CourseCreator{
		courseRepository: r,
	}
}
