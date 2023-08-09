package domain

import (
	shared "github.com/dasalgadoc/clean-architecture-go/shared/event"
)

type Course struct {
	shared.AggregateRoot
	id   CourseId
	name CourseName
}

func (c *Course) Id() string {
	return c.id.Value()
}

func (c *Course) Name() string {
	return c.name.Value()
}

func CreateCourse(name string) (*Course, error) {
	id, err := NewCourseId()
	if err != nil {
		return nil, err
	}
	studentName, err := NewCourseName(name)
	if err != nil {
		return nil, err
	}

	course := &Course{
		id:   *id,
		name: *studentName,
	}

	course.AggregateRoot.Record(NewCourseCreatedEvent(course.Id(), course.Name()))

	return course, nil
}

func NewCourseWithId(id string, name string) (*Course, error) {
	courseId, err := NewCourseIdFromString(id)
	if err != nil {
		return nil, err
	}
	courseName, err := NewCourseName(name)
	if err != nil {
		return nil, err
	}

	return &Course{
		id:   *courseId,
		name: *courseName,
	}, nil
}
