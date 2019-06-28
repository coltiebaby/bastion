package replay

import (
	"fmt"
	"net/http"

	"github.com/coltiebaby/go-lcu/clients"
	cu "github.com/coltiebaby/go-lcu/clients/clientutil"
)

type ReplayClient struct {
	Port string
}

func NewReplayClient() (*ReplayClient, error) {
	return &ReplayClient{Port: DEFAULT_PORT}, nil
}

func (c *ReplayClient) NewRequest(req_type, uri string, form []byte) (*http.Request, error) {
	rawUrl := fmt.Sprintf("https://127.0.0.1:%s%s", c.Port, uri)

	req, err := clients.DefaultNewRequest(req_type, rawUrl, form)
	if err != nil {
		return req, err
	}

	return req, nil
}

func (c *ReplayClient) Get(uri string) (*http.Response, error) {
	req, err := c.NewRequest("GET", uri, nil)
	if err != nil {
		return &http.Response{}, err
	}

	return cu.HttpClient.Do(req)
}

func (c *ReplayClient) Post(uri string, data []byte) (*http.Response, error) {
	req, err := c.NewRequest("POST", uri, data)
	if err != nil {
		return &http.Response{}, err
	}

	return cu.HttpClient.Do(req)
}

const (
	DEFAULT_PORT string = "2999"
)
