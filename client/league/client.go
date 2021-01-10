package league

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"

	"github.com/coltiebaby/bastion/client"
	cu "github.com/coltiebaby/bastion/client/clientutil"
)

type Client struct {
	token string
	Port  string
	Path  string
}

// Create From Unix
//
// Creates a new client from an already open league of legends client using commands
// that are related to a unix based system
func CreateFromUnix() (client.Client, error) {
	some_byes, err := exec.Command("ps", "-A").Output()
	if err != nil {
		return &Client{}, NotRunningErr
	}

	cmd := exec.Command("grep", "ClientUx")
	// Mimic "piping" data from a cmd
	cmd.Stdin = bytes.NewReader(some_byes)

	output, err := cmd.Output()
	if err != nil {
		return &Client{}, NotRunningErr
	}

	return newClient(output)
}

// Create From Windows
//
// Creates a new client from an already open league of legends client using commands
// that are related to a windows based system
func CreateFromWindows() (client.Client, error) {
	cmd := []string{
		`process`,
		`where`,
		`name="ClientUx.exe"`,
		`get`,
		`Caption,Processid,Commandline`,
	}

	output, err := exec.Command(`wmic`, cmd...).Output()
	if err != nil {
		return &Client{}, NotRunningErr
	}

	return newClient(output)
}

// Both operating systems produce an output where we can find the important pieces for Client
func newClient(output []byte) (client.Client, error) {
	ports := regexp.MustCompile(`--app-port=([0-9]*)`).FindAllSubmatch(output, 1)
	paths := regexp.MustCompile(`--install-directory=([\w//-_]*)`).FindAllSubmatch(output, 1)
	tokens := regexp.MustCompile(`--remoting-auth-token=([\w-_]*)`).FindAllSubmatch(output, 1)

	if len(ports) < 0 && len(tokens) < 0 {
		return &Client{}, NotRunningErr
	}

	port := string(ports[0][1])
	token := string(tokens[0][1])
	path := string(paths[0][1])

	return &Client{token: token, Port: port, Path: path}, nil
}

// URL returns a url.URL that you can edit further.
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

	req.SetBasicAuth(`riot`, c.token)
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

var (
	DownloadFailedErr error = fmt.Errorf("Failed to download file.")
	NotRunningErr     error = errors.New("League of legends is not currently running!")
)
