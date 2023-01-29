package internal

import (
	"github.com/pkg/errors"
	"go.uber.org/config"
	"os"
	"strings"

	. "github.com/reifiedbeans/project-zombot/internal/discord"
)

type Config struct {
	Discord struct {
		ApplicationID Snowflake `yaml:"applicationId"`
		PublicKey     string    `yaml:"publicKey"`
		Bot           struct {
			Token string
			Port  string
		}
		AllowedGuilds []Snowflake `yaml:"allowedGuilds"`
	}
	RCON struct {
		Host     string
		Port     string
		Password string
	}
	Game struct {
		Password string
	}
}

// The list of config files is ordered in priority, with the highest last.
// This is due to how config.NewYAML applies config options.
var files = []string{
	"/usr/local/share/pzombot/config.yaml",
	"config.yaml",
}

const defaults = `
discord:
  applicationId: ${DISCORD_APPLICATION_ID}
  publicKey: ${DISCORD_PUBLIC_KEY}
  bot:
    token: ${DISCORD_BOT_TOKEN}
    port: ${DISCORD_BOT_PORT:9268}
  allowedGuilds:
    - ${DISCORD_GUILD_ID}
rcon:
  host: ${RCON_HOST:0.0.0.0}
  port: ${RCON_PORT:27015}
  password: ${RCON_PASSWORD}
game:
  password: ${GAME_PASSWORD}
`

func NewConfig() *Config {
	options := []config.YAMLOption{
		config.Expand(os.LookupEnv),
		config.Source(strings.NewReader(defaults)),
	}

	for _, file := range files {
		if _, err := os.Stat(file); err != nil {
			continue
		}
		options = append(options, config.File(file))
	}

	provider, err := config.NewYAML(options...)
	if err != nil {
		err = errors.Wrap(err, "Failed to construct new config provider")
		panic(err)
	}

	var c Config
	if err = provider.Get(config.Root).Populate(&c); err != nil {
		err = errors.Wrap(err, "Failed to populate config")
		panic(err)
	}

	return &c
}
