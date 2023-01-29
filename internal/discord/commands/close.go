package commands

import (
	tempest "github.com/Amatsagu/Tempest"
	"go.uber.org/zap"

	. "github.com/reifiedbeans/project-zombot/internal/zomboid/rcon"
)

func NewCloseCommand(rcon *Rcon, log *zap.SugaredLogger) *Cmd {
	return &Cmd{
		Command: tempest.Command{
			Name:        "close",
			Description: "Stop allowing players to connect to the server",
			Type:        tempest.COMMAND_CHAT_INPUT,
			SlashCommandHandler: func(itx tempest.CommandInteraction) {
				it := CommandInteraction{
					CommandInteraction: itx,
				}

				if err := it.Defer(); err != nil {
					return
				}

				if err := rcon.CloseServer(); err != nil {
					log.Errorw("Failed to close server", zap.Error(err))
					_ = it.EditReply("Failed to close server. Contact an admin for help.")
					return
				}

				_ = it.EditReply("Server is now closed. Make sure all players have logged off.")
			},
		},
	}
}
