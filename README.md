# Project Zombot

A simple Discord bot for controlling and querying a Project Zomboid multiplayer server.

## Usage

This bot is self-hosted, as it needs to interact with your Project Zomboid server. You can follow these guides to get started:
- [Creating a Discord application](docs/creating-discord-application.md)
- [Deploying the bot to your game server](docs/deploying-to-server.md)
- [Configuration options](docs/configuration-options.md)

### Commands

Here is the list of slash commands currently supported by this bot.

| Name       | Description                                                                                                           |
|------------|-----------------------------------------------------------------------------------------------------------------------|
| `/players` | Show the current list of players logged on to the server.                                                             |
| `/open`    | Allow players to connect to the server. This is accomplished by setting the game password to a configured value.      |
| `/close`   | Stop allowing players to connect to the server. This is accomplished by changing the game password to a random value. |

## License

Licensed under the [MIT License](LICENSE).
