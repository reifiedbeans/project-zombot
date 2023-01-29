# Deploying the bot to your game server

This guide assumes that you have followed the [PZWiki instructions for running a dedicated server][dedicated-server] using systemd on Linux.

## Downloading the bot

Download the latest release from [GitHub][releases]. If none of the pre-compiled binaries match your operating system and architecture, follow the directions to [build from source](building-from-source.md). On a Linux system, copy the binary to `/usr/local/bin/pzombot`.

## Creating a run-as user

Create a new user to run the bot. This helps separate the bot from your game server for security.

```shell
adduser --disabled-password --gecos '' pzombot
```

## Configuring the bot

Create a `config.yaml` file in the home directory of the run-as user and paste the following into the file.

```yaml
discord:
  applicationId: "APPLICATION_ID_HERE"
  publicKey: "PUBLIC_KEY_HERE"
  bot:
    token: "Bot TOKEN_HERE"
  allowedGuilds:
    - "YOUR_GUILD_ID_HERE"
rcon:
  password: "RCON_PASSWORD_HERE"
game:
  password: "GAME_PASSWORD_HERE"
```

Fill in the `applicationId`, `publicKey`, and `token` using the values provided to you by Discord when you created your bot application. If you don't have a bot application yet, you can follow the [guide to create one](creating-discord-application.md). Make sure to prefix the bot token with `Bot`. Also add your server's ID to the `allowedGuilds`. Discord has [instructions][guild-id] on how to find your server's ID if you're not sure how.

The bot relies on the RCON protocol used by the Project Zomboid server in order to execute commands. Make sure you've set the `RCONPassword` in your [server settings][server-settings] and then copy it into the config file above.

Also copy your server's game password to the config file so that the bot can reset the password when the [`/open` command][open-command] is called.

There are other [configuration options](configuration-options.md) that you can add as well. If you'd rather not have a home directory for the bot, you can instead put the config file at `/usr/local/share/pzombot/config.yaml` or use the environment variables listed in the [configuration options](configuration-options.md).

## Running the bot

Create a new systemd unit file at `/usr/lib/systemd/system/zombot.service` and paste the following into the file.

```
[Unit]
Description=Project Zombot (Discord Bot)
Requisite=zomboid.service
After=zomboid.service

[Service]
PrivateTmp=true
Type=simple
User=pzombot
WorkingDirectory=/home/pzombot
ExecStart=/usr/local/bin/pzombot

[Install]
WantedBy=multi-user.target
```

Next, make sure systemd sleeps for some time after starting your server using the `ExecStartPost` option in your `zomboid.service` unit file.

```
[Service]
...
# Add this line
ExecStartPost=/bin/sleep 60
```

This ensures that the Project Zomboid server is fully started and RCON is running before the bot starts up.

Now you can run both the Zomboid server and the bot.

```shell
systemctl start zomboid zombot
```

If you need to see the logs for the bot, you can view them with `journalctl`.

```shell
journalctl -u zombot.service -f
```

## Setting up HTTPS

Discord requires that applications using Interactions communicate over HTTPS. The easiest way to set this up is using nginx with certbot.

DigitalOcean has a good guide for [setting up nginx][nginx-guide] if you're running your server on Ubuntu.

Before setting up `certbot`, create a new server block for the bot. The following is a simple configuration that works well. Create a new file at `/etc/nginx/sites-available/pzombot` and paste the following into the file.

```
server {
    server_name YOUR_DOMAIN_HERE;

    location / {
        proxy_pass http://localhost:9268/;
    }
}
```

If you modified the port the bot runs on, make sure to change it here as well.

Then you can enable the server configuration, test the config, and restart `nginx`.

```shell
ln -s /etc/nginx/sites-available/pzombot /etc/nginx/sites-enabled/
nginx -t
systemctl restart nginx
```

Next, set up certbot to get HTTPS working. DigitalOcean also has a guide for [setting up certbot with nginx][certbot-guide].

## Adding the bot endpoint

Once certbot is set up, you can add the bot URL to your Discord application in the [Discord developer portal][developer-portal]. Add your domain name, prefixed with `https://`, to the **Interactions Endpoint URL**.

---

You should be able to run commands in your Discord server now!

[dedicated-server]: https://pzwiki.net/wiki/Dedicated_Server
[releases]: https://github.com/reifiedbeans/project-zombot/releases
[guild-id]: https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID
[server-settings]: https://pzwiki.net/wiki/Server_Settings
[open-command]: /internal/discord/commands/open.go
[nginx-guide]: https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-22-04#step-5-setting-up-server-blocks-recommended
[certbot-guide]: https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-22-04#step-2-confirming-nginx-s-configuration
[developer-portal]: https://discord.com/developers
