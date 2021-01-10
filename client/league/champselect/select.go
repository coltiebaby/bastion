package champselect

import (
	"fmt"
	"net/http"

	"github.com/coltiebaby/bastion/client"
	cu "github.com/coltiebaby/bastion/client/clientutil"
)

type ChampSelect struct {
	lc client.Client
}

const fmtstr = "/lol-champ-select/v1/%s"

func createUri(uri string) string {
	return fmt.Sprintf(fmtstr, uri)
}

func New(lc client.Client) *ChampSelect {
	return &ChampSelect{
		lc: lc,
	}
}

func (cs *ChampSelect) Session() (*http.Response, error) {
	u, err := cs.lc.URL(createUri("session"))
	if err != nil {
		return nil, err
	}

	return cs.lc.Get(u)
}

func (cs *ChampSelect) PickableSkins() (*http.Response, error) {
	u, err := cs.lc.URL(createUri("session/my-selection"))
	if err != nil {
		return nil, err
	}

	req, _ := cs.lc.NewRequest("PATCH", u, []byte{})
	return cu.HttpClient.Do(req)

}
