# Bastion

interact with replay, live game, or the client api for league of legends

## Getting Started

### Game Config
The `game.cfg` file is located in

- MacOS
`/Applications/League\ of\ Legends.app/Contents/LoL/Config/game.cfg`

- Windows
`C:\Riot Games\League Of Legends\Config\game.cfg`

*Note: You must have played a game once before its created.*

### Enabling TLS
Attempt to look at two places when your app first launches.
1. Looks for the environmental `LCU_SSH` to be set to an absolute path
2. Checks if its in a default spot
-- windows `C:\Documents\riotgames.pem`
-- unix/darwin `~/Documents/riotgames.pem`

The pem file is located in [Riots Documentation](https://developer.riotgames.com/replay-apis.html) under root certificate link.

## Useful Documentation
Some information that will help you with your journey
- [Riots API Documentation](https://developer.riotgames.com/)
- [Riot's Discord](https://discord.gg/RiotGamesAPI)
- [vivide.re](http://lcu.vivide.re/)

License
----

MIT
