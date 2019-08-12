package replays

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coltiebaby/go-lcu/clients"
)

type Replay struct {
	client  clients.Client
	MatchId string
}

func NewReplay(client clients.Client, matchId string) Replay {
	return Replay{
		client:  client,
		MatchId: matchId,
	}
}

func (r Replay) postDownload(ctx *clients.Context, endpoint string) (err error) {
	uri := fmt.Sprintf("/lol-replays/v1/rofls/%s/download"+endpoint, r.MatchId)

	data, err := json.Marshal(ctx)
	if err != nil {
		return resp, err
	}

	if resp, err = r.client.Post(uri, data); err != nil && resp.StatusCode != 204 {
		err = fmt.Errorf("Failed to download replay! Error: %s", err)
	}
	return err
}

func (r Replay) Path() (path string, err error) {
	uri := `/lol-replays/v1/rofls/path`

	resp, err := r.client.Get(uri)
	if err != nil && resp.StatusCode != 204 {
		err = fmt.Errorf("Failed to download replay! Error: %s", err)
		return path, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&path); err != nil {
		err = fmt.Errorf("Could not decode replay path! Error: %s", err)
	}
	return path, err
}

func (r Replay) Download() error {
	ctx := clients.NewContext()
	ctx.AddComponent(`contextData`, MatchHistoryButton)

	return r.postDownload(ctx, "")
}

func (r Replay) DownloadGraceful(client clients.Client) error {
	ctx := clients.NewContext()
	ctx.AddComponent(`contextData`, MatchHistoryButton)

	return r.postDownload(ctx, "/graceful")
}

// Buttons
var (
	MatchHistoryButton = clients.NewComponent(`replay-button_match-history`)
)
