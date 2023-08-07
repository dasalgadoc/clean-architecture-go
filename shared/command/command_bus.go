package command

import "context"

type CommandBus interface {
	Dispatch(ctx context.Context, command Command) error
	Register(Type, Handler)
}
