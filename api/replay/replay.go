package replay

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/coltiebaby/bastion/api"
	"github.com/coltiebaby/bastion/client"
	"github.com/coltiebaby/bastion/components"
)

type Replay struct {
	client  client.Client
	MatchId string
}

func New(client client.Client, matchId string) Replay {
	return Replay{
		client:  client,
		MatchId: matchId,
	}
}

func (r Replay) NewURL(endpoint string) (url.URL, error) {
	req := api.Request{
		Domain:  "lol-replays",
		Version: "v1",
		Uri:     endpoint,
	}

	return r.client.URL(req.String())
}

// Checks the client for the configuration
func (r Replay) GetConfiguration() (c Config, err error) {
	u, err := r.NewURL(fmt.Sprintf(`configuration`))
	if err != nil {
		return c, err
	}

	if resp, err := r.client.Get(u); err != nil {
		err = json.NewDecoder(resp.Body).Decode(&c)
	}

	return c, err
}

func (r Replay) GetMetadata() (m Meta, err error) {
	u, err := r.NewURL(fmt.Sprintf(`metadata/%s`, r.MatchId))
	if err != nil {
		return m, err
	}

	if resp, err := r.client.Get(u); err != nil {
		err = json.NewDecoder(resp.Body).Decode(&m)
	}

	return m, err
}

// Returns the current replay directory set
func (r Replay) path(endpoint string) (path string, err error) {
	u, err := r.NewURL(fmt.Sprintf(endpoint))
	if err != nil {
		return path, err
	}

	if resp, err := r.client.Get(u); err != nil {
		err = json.NewDecoder(resp.Body).Decode(&path)
	}

	return path, err
}

// Gets the current path to the replay folder
func (r Replay) Path() (path string, err error) {
	path, err = r.path(`rofls/path`)
	return path, err
}

// Gets the default path to the replay folder
func (r Replay) PathDefault() (path string, err error) {
	path, err = r.path(`rofls/path/default`)
	return path, err
}

func (r Replay) Scan() (err error) {
	u, err := r.NewURL("/rofls/scan")
	if err != nil {
		return err
	}

	// Looks for a 204
	_, err = r.client.Get(u)

	return err
}

func (r Replay) postDownload(endpoint string) (err error) {
	u, err := r.NewURL(fmt.Sprintf("rofls/%s/%s", r.MatchId, endpoint))
	if err != nil {
		return err
	}

	ctx := components.NewContext()
	ctx.AddComponent(`contextData`, MatchHistoryButton)

	data, err := json.Marshal(ctx)
	if err != nil {
		return err
	}

	_, err = r.client.Post(u, data)
	return err
}

func (r Replay) Download() (err error) {
	return r.postDownload("download")
}

func (r Replay) DownloadGraceful() error {
	return r.postDownload("download/graceful")
}

func (r Replay) Watch() (err error) {
	u, err := r.NewURL(fmt.Sprintf("rofls/%s/watch", r.MatchId))
	if err != nil {
		return err
	}

	_, err = r.client.Post(u, []byte{})

	return err
}

// Buttons
var (
	MatchHistoryButton = components.NewComponent(`replay-button_match-history`)
)
