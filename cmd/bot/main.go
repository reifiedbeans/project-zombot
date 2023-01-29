package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	. "github.com/reifiedbeans/project-zombot/internal"
	. "github.com/reifiedbeans/project-zombot/internal/discord/bot"
	. "github.com/reifiedbeans/project-zombot/internal/discord/commands"
	. "github.com/reifiedbeans/project-zombot/internal/logging"
	. "github.com/reifiedbeans/project-zombot/internal/zomboid/rcon"
)

func main() {
	// Dependency injection wiring
	fx.New(
		fx.Provide(
			NewConfig,
			NewLogger,
			NewRcon,
			NewBot,
			asCommand(NewPlayersCommand),
			asCommand(NewOpenCommand),
			asCommand(NewCloseCommand),
		),
		fx.WithLogger(func(log *zap.SugaredLogger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log.Desugar()}
		}),
		fx.Invoke(func(*Bot) {}),
	).Run()
}

func asCommand(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Command)),
		fx.ResultTags(`group:"commands"`),
	)
}
