package replay

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coltiebaby/bastion/clients"
	"github.com/coltiebaby/bastion/components"
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

func (r Replay) fmtUri(endpoint string, opts ...string) string {
	base := fmt.Sprintf(`/lol-replays/v1/%s`, endpoint)
	return fmt.Sprintf(base, strings.Join(opts, `/`))
}

// Checks the client for the configuration
func (r Replay) GetConfiguration() (c Config, err error) {
	uri := r.fmtUri(`configuration`)
	if resp, err := r.client.Get(uri); err != nil {
		err = json.NewDecoder(resp.Body).Decode(&c)
	}

	return c, err
}

func (r Replay) GetMetadata() (m Meta, err error) {
	uri := r.fmtUri(`/metadata/%s`, r.MatchId)

	if resp, err := r.client.Get(uri); err != nil {
		err = json.NewDecoder(resp.Body).Decode(&m)
	}

	return m, err
}

// Returns the current replay directory set
func (r Replay) path(endpoint string) (path string, err error) {
	uri := r.fmtUri(endpoint)

	if resp, err := r.client.Get(uri); err != nil {
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
	uri := fmt.Sprintf("rofls/scan")
	// Looks for a 204
	_, err = r.client.Get(uri)

	return err
}

func (r Replay) postDownload(endpoint string) (err error) {
	uri := r.fmtUri("rofls/%s"+endpoint, r.MatchId)

	ctx := components.NewContext()
	ctx.AddComponent(`contextData`, MatchHistoryButton)

	data, err := json.Marshal(ctx)
	if err != nil {
		return err
	}

	_, err = r.client.Post(uri, data)
	return err
}

func (r Replay) Download() (err error) {
	return r.postDownload("/download")
}

func (r Replay) DownloadGraceful() error {
	return r.postDownload("/download/graceful")
}

func (r Replay) Watch() (err error) {
	uri := r.fmtUri("rofls/%s/watch", r.MatchId)

	_, err = r.client.Post(uri, []byte{})

	return err
}

// Buttons
var (
	MatchHistoryButton = components.NewComponent(`replay-button_match-history`)
)
