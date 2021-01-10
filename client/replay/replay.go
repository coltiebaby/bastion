package replay

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/coltiebaby/bastion/client"
	cu "github.com/coltiebaby/bastion/client/clientutil"
)

type ReplayClient struct {
	Port string
}

func NewReplayClient() (*ReplayClient, error) {
	return &ReplayClient{Port: DEFAULT_PORT}, nil
}

func (c *LeagueClient) URL(uri string) (url.URL, error) {
	return url.Parse(fmt.Sprintf("https://127.0.0.1:%s%s", c.Port, uri))
}

func (c *ReplayClient) NewRequest(req_type, u url.URL, form []byte) (*http.Request, error) {

	req, err := client.DefaultNewRequest(req_type, u.String(), form)
	if err != nil {
		return req, err
	}

	return req, nil

}

func (c *ReplayClient) Get(u url.URL) (*http.Response, error) {
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return &http.Response{}, err
	}

	return cu.HttpClient.Do(req)
}

func (c *ReplayClient) Post(u url.URL, data []byte) (*http.Response, error) {
	req, err := c.NewRequest("POST", u, data)
	if err != nil {
		return &http.Response{}, err
	}

	return cu.HttpClient.Do(req)
}

const (
	DEFAULT_PORT string = "2999"
)
