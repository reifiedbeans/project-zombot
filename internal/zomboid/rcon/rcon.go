package rcon

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/gorcon/rcon"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"go.uber.org/zap"

	. "github.com/reifiedbeans/project-zombot/internal"
)

type Rcon struct {
	rcon.Conn

	log    *zap.SugaredLogger
	config *Config
}

type RconParams struct {
	fx.In

	Log    *zap.SugaredLogger
	Config *Config
}

func NewRcon(lc fx.Lifecycle, p RconParams) (r *Rcon) {
	r = &Rcon{
		log:    p.Log,
		config: p.Config,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := fmt.Sprintf("%s:%s", p.Config.RCON.Host, p.Config.RCON.Port)
			conn, err := rcon.Dial(addr, p.Config.RCON.Password)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("Failed to connect to '%s' over RCON with given password", addr))
			}

			r.Conn = *conn
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := r.Close(); err != nil {
				return errors.Wrap(err, "Failed to disconnect from RCON")
			}
			return nil
		},
	})
	return r
}

func (r *Rcon) GetPlayers() (players []string, err error) {
	var playerRegex = regexp.MustCompile("^-(.+)$")

	output, err := r.Execute("players")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to execute 'players' command")
	}
	output = strings.Trim(output, "\n")

	// Slice ignores first line since it's a header
	lines := strings.Split(output, "\n")[1:]
	for _, line := range lines {
		groups := playerRegex.FindStringSubmatch(line)
		if len(groups) != 2 {
			return nil, errors.New(fmt.Sprintf("Could not parse player from '%s'", line))
		}
		players = append(players, groups[1])
	}
	return players, nil
}

func (r *Rcon) OpenServer() error {
	if err := r.setPassword(r.config.Game.Password); err != nil {
		return errors.Wrap(err, "Failed to open server")
	}
	return nil
}

func (r *Rcon) CloseServer() error {
	pw := strings.ReplaceAll(uuid.New().String(), "-", "")
	if err := r.setPassword(pw); err != nil {
		return errors.Wrap(err, "Failed to close server")
	}
	return nil
}

func (r *Rcon) setPassword(pw string) error {
	var outputRegex = regexp.MustCompile("^Option : Password is now : (.+)$")

	r.log.Infow(
		"Setting game server password",
		zap.String("newPassword", pw),
	)
	output, err := r.Execute(fmt.Sprintf("changeoption Password \"%s\"", pw))
	if err != nil {
		return errors.Wrap(err, "Failed to set password using 'changeoption' command")
	}
	output = strings.Trim(output, "\n")

	groups := outputRegex.FindStringSubmatch(output)
	if groups[1] != pw {
		return errors.New("Server did not successfully set new password")
	}

	return nil
}
