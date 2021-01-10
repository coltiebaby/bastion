package champselect

type Action struct {
	ActorCellID  int    `json:"actorCellId"`
	ChampionID   int    `json:"championId"`
	Completed    bool   `json:"completed"`
	ID           int    `json:"id"`
	IsAllyAction bool   `json:"isAllyAction"`
	IsInProgress bool   `json:"isInProgress"`
	PickTurn     int    `json:"pickTurn"`
	Type         string `json:"type"`
}

type Bans struct {
	MyTeamBans    []int `json:"myTeamBans"`
	NumBans       int   `json:"numBans"`
	TheirTeamBans []int `json:"theirTeamBans"`
}

// TODO: Hiding this; Not allowed to access chats
type ChatDetail struct {
	chatRoomName     string `json:"chatRoomName"`
	chatRoomPassword string `json:"chatRoomPassword"`
}

type EntitledFeatureState struct {
	AdditionalRerolls int           `json:"additionalRerolls"`
	UnlockedSkinIds   []interface{} `json:"unlockedSkinIds"`
}

type Player struct {
	AssignedPosition    string `json:"assignedPosition"`
	CellID              int    `json:"cellId"`
	ChampionID          int    `json:"championId"`
	ChampionPickIntent  int    `json:"championPickIntent"`
	EntitledFeatureType string `json:"entitledFeatureType"`
	SelectedSkinID      int    `json:"selectedSkinId"`
	Spell1ID            int    `json:"spell1Id"`
	Spell2ID            int    `json:"spell2Id"`
	SummonerID          int    `json:"summonerId"`
	Team                int    `json:"team"`
	WardSkinID          int    `json:"wardSkinId"`
}

type Timer struct {
	AdjustedTimeLeftInPhase int    `json:"adjustedTimeLeftInPhase"`
	InternalNowInEpochMs    int64  `json:"internalNowInEpochMs"`
	IsInfinite              bool   `json:"isInfinite"`
	Phase                   string `json:"phase"`
	TotalTimeInPhase        int    `json:"totalTimeInPhase"`
}

type Session struct {
	Actions              [][]Action           `json:"actions"`
	AllowBattleBoost     bool                 `json:"allowBattleBoost"`
	AllowDuplicatePicks  bool                 `json:"allowDuplicatePicks"`
	AllowLockedEvents    bool                 `json:"allowLockedEvents"`
	AllowRerolling       bool                 `json:"allowRerolling"`
	AllowSkinSelection   bool                 `json:"allowSkinSelection"`
	Bans                 Bans                 `json:"bans"`
	BenchChampionIds     []int                `json:"benchChampionIds"`
	BenchEnabled         bool                 `json:"benchEnabled"`
	BoostableSkinCount   int                  `json:"boostableSkinCount"`
	ChatDetails          ChatDetail           `json:"chatDetails"`
	Counter              int                  `json:"counter"`
	EntitledFeatureState EntitledFeatureState `json:"entitledFeatureState"`
	GameID               int                  `json:"gameId"`
	HasSimultaneousBans  bool                 `json:"hasSimultaneousBans"`
	HasSimultaneousPicks bool                 `json:"hasSimultaneousPicks"`
	IsCustomGame         bool                 `json:"isCustomGame"`
	IsSpectating         bool                 `json:"isSpectating"`
	LocalPlayerCellID    int                  `json:"localPlayerCellId"`
	LockedEventIndex     int                  `json:"lockedEventIndex"`
	MyTeam               []Player             `json:"myTeam"`
	RerollsRemaining     int                  `json:"rerollsRemaining"`
	SkipChampionSelect   bool                 `json:"skipChampionSelect"`
	TheirTeam            []Player             `json:"theirTeam"`
	Timer                Timer                `json:"timer"`
	Trades               []interface{}        `json:"trades"`
}
