package command

import "context"

type Type string

type Command interface {
	Type() Type
}

type Handler interface {
	Handle(ctx context.Context, command Command) error
}
