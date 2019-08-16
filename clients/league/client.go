package league

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"

	"github.com/coltiebaby/go-lcu/clients"
	cu "github.com/coltiebaby/go-lcu/clients/clientutil"
)

type LeagueClient struct {
	token string
	Port  string
	Path  string
}

// Create client from process information
func CreateFromUnix() (clients.Client, error) {
	some_byes, err := exec.Command("ps", "-A").Output()
	if err != nil {
		return &LeagueClient{}, NOT_RUNNING_ERR
	}

	cmd := exec.Command("grep", "LeagueClientUx")
	// Mimic "piping" data from a cmd
	cmd.Stdin = bytes.NewReader(some_byes)

	output, err := cmd.Output()
	if err != nil {
		return &LeagueClient{}, NOT_RUNNING_ERR
	}

	ports := regexp.MustCompile(`--app-port=([0-9]*)`).FindAllSubmatch(output, 1)
	paths := regexp.MustCompile(`--install-directory=([\w//-_]*)`).FindAllSubmatch(output, 1)
	tokens := regexp.MustCompile(`--remoting-auth-token=([\w-_]*)`).FindAllSubmatch(output, 1)

	if len(ports) < 0 && len(tokens) < 0 {
		return &LeagueClient{}, NOT_RUNNING_ERR
	}

	port := string(ports[0][1])
	token := string(tokens[0][1])
	path := string(paths[0][1])

	return &LeagueClient{token: token, Port: port, Path: path}, nil
}

func (c *LeagueClient) NewRequest(req_type, uri string, form []byte) (*http.Request, error) {
	rawUrl := fmt.Sprintf("https://127.0.0.1:%s%s", c.Port, uri)

	req, err := clients.DefaultNewRequest(req_type, rawUrl, form)
	if err != nil {
		return req, err
	}
	req.SetBasicAuth(`riot`, c.token)

	return req, nil
}

func (c *LeagueClient) Get(uri string) (*http.Response, error) {
	req, err := c.NewRequest("GET", uri, nil)
	if err != nil {
		return &http.Response{}, err
	}

	return cu.HttpClient.Do(req)
}

func (c *LeagueClient) Post(uri string, data []byte) (*http.Response, error) {
	req, err := c.NewRequest("POST", uri, data)
	if err != nil {
		return &http.Response{}, err
	}

	return cu.HttpClient.Do(req)
}

var (
	DownloadFailedErr error = fmt.Errorf("Failed to download file.")
	NOT_RUNNING_ERR   error = errors.New("League of legends is not currently running!")
)
