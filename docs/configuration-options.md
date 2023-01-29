# Configuration options

The following is an empty config file with all the available options. All options can also be configured using environment variables.

```yaml
discord:
  applicationId: "" # or using $DISCORD_APPLICATION_ID
  publicKey: "" # or using $DISCORD_PUBLIC_KEY
  bot:
    token: "" # or using $DISCORD_BOT_TOKEN
    port: "" # or using $DISCORD_BOT_PORT, default: 9268
  allowedGuilds:
    - "" # or using $DISCORD_GUILD_ID (only one via environment variable)
rcon:
  host: "" # or using $RCON_HOST, default: 0.0.0.0
  port: "" # or using $RCON_PORT, default: 27015
  password: "" # or using $RCON_PASSWORD
game:
  password: "" # or using $GAME_PASSWORD
```

Below are the valid locations for the config file. They are listed in the order of priority (high to low).
- `./config.yaml` (same location as the bot binary)
- `/usr/local/share/pzombot/config.yaml`
