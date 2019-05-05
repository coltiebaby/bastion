# GO-LCU

go-lcu is a library that helps you get connected to your league client with the replay api added in for good measure.

# Getting Started
##### Enabling TLS
Attempt to look at two places when your app first launches.
1. Looks for the environmental `LCU_SSH` to be set to an absolute path
2. Checks if its in a default spot
-- windows `C:\\Documents\riotgames.pem`
-- unix/darwin `~/Documents/riotgames.pem`

The pem file is located at riot's [documentation](https://developer.riotgames.com/replay-apis.html) under root certificate link.

###### Useful Documentation
Some information that will help you with your journey
- [Riots API Documentation](https://developer.riotgames.com/)
- [Riot's Discord](https://discord.gg/RiotGamesAPI)
- [vivide.re](http://lcu.vivide.re/)

License
----

MIT
