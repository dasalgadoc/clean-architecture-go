package domain

import "errors"

type CourseName struct {
	value string
}

func (c *CourseName) Value() string {
	return c.value
}

func NewCourseName(value string) (*CourseName, error) {
	err := courseNameNotEmpty(value)
	if err != nil {
		return nil, err
	}

	return &CourseName{value: value}, nil
}

func courseNameNotEmpty(value string) error {
	if value == "" {
		return errors.New("course name is empty")
	}
	return nil
}
