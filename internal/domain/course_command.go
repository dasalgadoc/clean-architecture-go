package domain

import "github.com/dasalgadoc/clean-architecture-go/shared/command"

const CourseCommandType command.Type = "command.course.create"

type CourseCommand struct {
	name string
}

func NewCourseCommand(name string) *CourseCommand {
	return &CourseCommand{
		name: name,
	}
}

func (c *CourseCommand) Type() command.Type {
	return CourseCommandType
}

func (c *CourseCommand) Name() string {
	return c.name
}
