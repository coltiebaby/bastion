// Holds the different types of data models we receive from the client
package replay

type Config struct {
	GameVersion                      string `json:"gameVersion"`
	IsInTournament                   bool   `json:"isInTournament"`
	IsLoggedIn                       bool   `json:"isLoggedIn"`
	IsPatching                       bool   `json:"isPatching"`
	IsPlayingGame                    bool   `json:"isPlayingGame"`
	IsPlayingReplay                  bool   `json:"isPlayingReplay"`
	IsReplaysEnabled                 bool   `json:"isReplaysEnabled"`
	IsReplaysForEndOfGameEnabled     bool   `json:"isReplaysForEndOfGameEnabled"`
	IsReplaysForMatchHistoryEnabled  bool   `json:"isReplaysForMatchHistoryEnabled"`
	MinServerVersion                 string `json:"minServerVersion"`
	MinutesUntilReplayConsideredLost int    `json:"minutesUntilReplayConsideredLost"`
}

type Meta struct {
	DownloadProgress int    `json:"downloadProgress"`
	GameID           int    `json:"gameId"`
	State            string `json:"state"`
}
