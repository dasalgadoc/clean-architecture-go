package memory

import (
	"context"
	"github.com/dasalgadoc/clean-architecture-go/shared/command"
	"log"
)

type CommandBus struct {
	handlers map[command.Type]command.Handler
}

func (c *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := c.handlers[cmd.Type()]
	if !ok {
		return nil
	}

	go func() {
		err := handler.Handle(ctx, cmd)
		if err != nil {
			log.Println(err)
		}
	}()

	// return handler.Handle(ctx, cmd) // Sync version

	return nil
}

func (c *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	c.handlers[cmdType] = handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}
