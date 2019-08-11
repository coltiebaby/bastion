package replays

import (
	"fmt"

	"github.com/coltiebaby/go-lcu/clients"
)

type Replay struct {
	MatchId string
}

func (r Replay) postDownload(ctx clients.Context, endpoint string) (resp *http.Response, err error) {
	uri := fmt.Sprintf("/lol-replays/v1/rofls/%s/download"+endpoint, r.MatchId)

	data, err := json.Marshal(ctx)
	if err != nil {
		return resp, err
	}

	if resp, err = clients.Post(uri, data); err != nil && resp.StatusCode != 204 {
		err = fmt.Error("Failed to download replay! Error: %s", err)
	}
	return resp, err
}

func (r Replay) Download(client clients.Client) error {
	ctx := clients.NewContext()
	ctx.AddComponent(`contextData`, MatchHistoryButton)

	resp, err := postDownload(ctx, "")
}

func (r Replay) DownloadGraceful(client clients.Client) error {
	ctx := clients.NewContext()
	ctx.AddComponent(`contextData`, MatchHistoryButton)

	resp, err := postDownload(ctx, "/graceful")
}

// Buttons
var (
	MatchHistoryButton = clients.NewComponent(`replay-button_match-history`)
)
