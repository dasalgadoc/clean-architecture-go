package domain

import shared "github.com/dasalgadoc/clean-architecture-go/shared/domain"

type CourseId struct {
	value shared.UUIDValueObject
}

func (c *CourseId) Value() string {
	return c.value.Value()
}

func NewCourseId() (*CourseId, error) {
	id, err := shared.NewUUIDValueObject()
	if err != nil {
		return nil, err
	}

	return &CourseId{value: *id}, nil
}

func NewCourseIdFromString(value string) (*CourseId, error) {
	id, err := shared.NewUUIDValueObjectFromString(value)
	if err != nil {
		return nil, err
	}

	return &CourseId{value: *id}, nil
}
