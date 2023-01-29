package bot

import (
	"context"
	"fmt"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"go.uber.org/zap"

	. "github.com/reifiedbeans/project-zombot/internal"
	. "github.com/reifiedbeans/project-zombot/internal/discord/commands"
)

type Bot struct {
	tempest.Client

	commands map[string]Command
}

type BotParams struct {
	fx.In

	Commands []Command `group:"commands"`
	Log      *zap.SugaredLogger
	Config   *Config
}

func NewBot(lc fx.Lifecycle, p BotParams) (b *Bot) {
	cmap := make(map[string]Command)
	for _, cmd := range p.Commands {
		cmap[cmd.Get().Name] = cmd
	}

	b = &Bot{
		commands: cmap,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := fmt.Sprintf("0.0.0.0:%s", p.Config.Discord.Bot.Port)

			b.Client = tempest.CreateClient(tempest.ClientOptions{
				ApplicationID: p.Config.Discord.ApplicationID.Snowflake,
				PublicKey:     p.Config.Discord.PublicKey,
				Token:         p.Config.Discord.Bot.Token,
				PreCommandExecutionHandler: func(itx tempest.CommandInteraction) *tempest.ResponseData {
					p.Log.Infow(
						"Starting execution of command",
						zap.String("command", itx.Data.Name),
						zap.String("actor", itx.Member.User.Tag()),
						zap.String("actorId", itx.Member.User.ID.String()),
					)
					return nil
				},
			})

			for _, cmd := range p.Commands {
				if err := b.Client.RegisterCommand(cmd.Get()); err != nil {
					return errors.Wrap(err, fmt.Sprintf("Failed to register command '%s'", cmd.Get().Name))
				}
			}

			var guilds []tempest.Snowflake
			for _, guild := range p.Config.Discord.AllowedGuilds {
				guilds = append(guilds, guild.Snowflake)
			}

			if err := b.Client.SyncCommands(guilds, nil, false); err != nil {
				return errors.Wrap(err, "Failed to synchronize commands with Discord")
			}

			go b.Client.ListenAndServe("/", addr)
			return nil
		},
	})
	return b
}
