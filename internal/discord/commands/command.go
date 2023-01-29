package commands

import (
	tempest "github.com/Amatsagu/Tempest"
	"go.uber.org/zap"
)

type CommandInteraction struct {
	tempest.CommandInteraction

	Log *zap.SugaredLogger
}

func (it *CommandInteraction) Defer() error {
	if err := it.CommandInteraction.Defer(false); err != nil {
		it.Log.Errorw("Could not defer reply", zap.Error(err))
		return err
	}
	return nil
}

func (it *CommandInteraction) EditReply(msg string) error {
	res := tempest.ResponseData{
		Content: msg,
	}
	if err := it.CommandInteraction.EditReply(res, false); err != nil {
		it.Log.Errorw("Could not edit reply", zap.Error(err))
		return err
	}
	return nil
}

type Command interface {
	Get() tempest.Command
	Testing() bool
}

type Cmd struct {
	tempest.Command

	TestCommand bool
}

func (c Cmd) Get() tempest.Command {
	return c.Command
}

func (c Cmd) Testing() bool {
	return c.TestCommand
}

// Compile-time check
var _ Command = (*Cmd)(nil)
