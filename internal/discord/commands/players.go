package commands

import (
	"fmt"
	"strings"

	tempest "github.com/Amatsagu/Tempest"
	"go.uber.org/zap"

	. "github.com/reifiedbeans/project-zombot/internal/zomboid/rcon"
)

func NewPlayersCommand(rcon *Rcon, log *zap.SugaredLogger) *Cmd {
	return &Cmd{
		Command: tempest.Command{
			Name:        "players",
			Description: "Show the current list of players logged on to the server",
			Type:        tempest.COMMAND_CHAT_INPUT,
			SlashCommandHandler: func(itx tempest.CommandInteraction) {
				it := CommandInteraction{
					CommandInteraction: itx,
					Log:                log,
				}

				if err := it.Defer(); err != nil {
					return
				}

				players, err := rcon.GetPlayers()
				if err != nil {
					msg := "Could not get players from server"
					log.Errorw(msg, zap.Error(err))

					_ = it.EditReply(msg)
					return
				}

				content := strings.Builder{}
				content.WriteString(fmt.Sprintf("**Players online (%d) :**", len(players)))

				for _, player := range players {
					content.WriteString("\n")
					content.WriteString(fmt.Sprintf("  - %s", player))
				}

				_ = it.EditReply(content.String())
			},
		},
	}
}
