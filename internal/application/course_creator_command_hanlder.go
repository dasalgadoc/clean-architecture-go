package application

import (
	"context"
	"errors"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"github.com/dasalgadoc/clean-architecture-go/shared/command"
)

type CourseCreatorCommandHandler struct {
	service CourseCreator
}

func NewCourseCreatorCommandHandler(service CourseCreator) CourseCreatorCommandHandler {
	return CourseCreatorCommandHandler{
		service: service,
	}
}

func (h *CourseCreatorCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	c, ok := cmd.(*domain.CourseCommand)
	if !ok {
		return errors.New("invalid command")
	}

	return h.service.Invoke(ctx, c.Name())
}
