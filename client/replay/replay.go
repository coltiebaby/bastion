package replay

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/coltiebaby/bastion/client"
	cu "github.com/coltiebaby/bastion/client/clientutil"
)

type Client struct {
	Port string
}

func New() (*Client, error) {
	return &Client{Port: DEFAULT_PORT}, nil
}

func (c *Client) URL(uri string) (u url.URL, err error) {
	urlp, err := url.Parse(fmt.Sprintf("https://127.0.0.1:%s%s", c.Port, uri))
	if err == nil {
		u = *urlp
	}

	return u, err
}

func (c *Client) NewRequest(req_type string, u url.URL, form []byte) (*http.Request, error) {

	req, err := client.DefaultNewRequest(req_type, u, form)
	if err != nil {
		return req, err
	}

	return req, nil

}

func (c *Client) Get(u url.URL) (*http.Response, error) {
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return &http.Response{}, err
	}

	return cu.HttpClient.Do(req)
}

func (c *Client) Post(u url.URL, data []byte) (*http.Response, error) {
	req, err := c.NewRequest("POST", u, data)
	if err != nil {
		return &http.Response{}, err
	}

	return cu.HttpClient.Do(req)
}

const (
	DEFAULT_PORT string = "2999"
)
