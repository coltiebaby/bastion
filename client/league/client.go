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

type LeagueClient struct {
	token string
	Port  string
	Path  string
}

func CreateFromUnix() (client.Client, error) {
	some_byes, err := exec.Command("ps", "-A").Output()
	if err != nil {
		return &LeagueClient{}, NotRunningErr
	}

	cmd := exec.Command("grep", "LeagueClientUx")
	// Mimic "piping" data from a cmd
	cmd.Stdin = bytes.NewReader(some_byes)

	output, err := cmd.Output()
	if err != nil {
		return &LeagueClient{}, NotRunningErr
	}

	return newClient(output)
}

func CreateFromWindows() (client.Client, error) {
	cmd := []string{
		`process`,
		`where`,
		`name="LeagueClientUx.exe"`,
		`get`,
		`Caption,Processid,Commandline`,
	}

	output, err := exec.Command(`wmic`, cmd...).Output()
	if err != nil {
		return &LeagueClient{}, NotRunningErr
	}

	return newClient(output)
}

func newClient(output []byte) (client.Client, error) {
	ports := regexp.MustCompile(`--app-port=([0-9]*)`).FindAllSubmatch(output, 1)
	paths := regexp.MustCompile(`--install-directory=([\w//-_]*)`).FindAllSubmatch(output, 1)
	tokens := regexp.MustCompile(`--remoting-auth-token=([\w-_]*)`).FindAllSubmatch(output, 1)

	if len(ports) < 0 && len(tokens) < 0 {
		return &LeagueClient{}, NotRunningErr
	}

	port := string(ports[0][1])
	token := string(tokens[0][1])
	path := string(paths[0][1])

	return &LeagueClient{token: token, Port: port, Path: path}, nil
}

func (c *LeagueClient) URL(uri string) (url.URL, error) {
	u, err := url.Parse(fmt.Sprintf("https://127.0.0.1:%s%s", c.Port, uri))
	return *u, err
}

func (c *LeagueClient) NewRequest(req_type string, u url.URL, form []byte) (*http.Request, error) {
	req, err := client.DefaultNewRequest(req_type, u, form)
	if err != nil {
		return req, err
	}

	req.SetBasicAuth(`riot`, c.token)
	return req, nil

}

func (c *LeagueClient) Get(u url.URL) (*http.Response, error) {
	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return &http.Response{}, err
	}

	return cu.HttpClient.Do(req)
}

func (c *LeagueClient) Post(u url.URL, data []byte) (*http.Response, error) {
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
