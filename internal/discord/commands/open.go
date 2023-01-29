package commands

import (
	tempest "github.com/Amatsagu/Tempest"
	"go.uber.org/zap"

	. "github.com/reifiedbeans/project-zombot/internal/zomboid/rcon"
)

func NewOpenCommand(rcon *Rcon, log *zap.SugaredLogger) *Cmd {
	return &Cmd{
		Command: tempest.Command{
			Name:        "open",
			Description: "Allow players to connect to the server",
			Type:        tempest.COMMAND_CHAT_INPUT,
			SlashCommandHandler: func(itx tempest.CommandInteraction) {
				it := CommandInteraction{
					CommandInteraction: itx,
					Log:                log,
				}

				if err := it.Defer(); err != nil {
					return
				}

				if err := rcon.OpenServer(); err != nil {
					log.Errorw("Failed to open server", zap.Error(err))
					_ = it.EditReply("Failed to open server. Contact an admin for help.")
					return
				}

				_ = it.EditReply("Server is now open. Have fun!")
			},
		},
	}
}
