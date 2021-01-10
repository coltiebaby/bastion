package builtin

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/coltiebaby/bastion/client"
)

type Builtin struct {
	lc client.Client
}

func New(lc client.Client) *Builtin {
	return &Builtin{
		lc: lc,
	}
}

type AsyncRequest struct {
	Token string
}

func NewAsyncRequest(token string) *AsyncRequest {
	return &AsyncRequest{
		Token: token,
	}
}

func (builtin *Builtin) async(uri string, req *AsyncRequest) (*http.Response, error) {
	if req.Token == "" {
		return nil, fmt.Errorf("requires token")
	}

	u, err := builtin.lc.URL(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uri: %w")
	}

	var v url.Values
	v.Set("asyncToken", req.Token)

	u.RawQuery = v.Encode()

	var data []byte
	return builtin.lc.Post(u, data)
}

// Cancels the asynchronous operation or removes its completion status.
func (builtin *Builtin) AsyncDelete(req *AsyncRequest) (*http.Response, error) {
	return builtin.async("/AsyncDelete", req)
}

// Retrieves the result of a completed asynchronous operation.
func (builtin *Builtin) AsyncResult(req *AsyncRequest) (*http.Response, error) {
	return builtin.async("/AsyncResult", req)
}

// Retrieves details on the current state of an asynchronous operation.
func (builtin *Builtin) AsyncStatus(req *AsyncRequest) (*http.Response, error) {
	return builtin.async("/AsyncStatus", req)
}

// Attempts to cancel an asynchronous operation
func (builtin *Builtin) AsyncCancel(req *AsyncRequest) (*http.Response, error) {
	return builtin.async("/Cancel", req)
}

// Closes the connection.
func (builtin *Builtin) Exit() (*http.Response, error) {
	u, err := builtin.lc.URL("/Exit")
	if err != nil {
		return nil, fmt.Errorf("failed to parse uri: %w")
	}

	var data []byte
	return builtin.lc.Post(u, data)
}

// With no arguments, returns a list of all available functions and types along with a short description
// If a function or type is specified, returns detailed information about it.
// TODO: look into formats. If empty it currently returns JSON
type HelpRequest struct {
	Target string // Function or Type name
	Format string // Full, Epytext, Brief, Console
}

// Returns information on available functions and types
func (builtin *Builtin) Help(req *HelpRequest) (*http.Response, error) {
	u, err := builtin.lc.URL("/Help")
	if err != nil {
		return nil, fmt.Errorf("failed to parse uri: %w")
	}

	var v url.Values

	if t := req.Target; t != "" {
		v.Set("target", t)
	}

	if f := req.Format; f != "" {
		v.Set("format", f)
	}

	u.RawQuery = v.Encode()

	var data []byte
	return builtin.lc.Post(u, data)
}

// Subscribes to a given event
func (builtin *Builtin) Subscribe() {

}

// Unsubscribes from a given event
// POST /Unsubscribe
func (builtin *Builtin) Unsubscribe() {

}

// With no arguments, returns the current output format being used.
// If a format is specified, switches the console output to that format.
type WebSocketFormatRequest struct {
	Format string // {JSON, YAML, MsgPack}
}

// Controls the console output format
func (builtin *Builtin) WebSocketFormat(req *WebSocketFormatRequest) (*http.Response, error) {
	u, err := builtin.lc.URL("/WebSocketFormat")
	if err != nil {
		return nil, fmt.Errorf("failed to parse uri: %w")
	}

	var v url.Values

	if f := req.Format; f != "" {
		v.Set("format", f)
	}

	u.RawQuery = v.Encode()

	var data []byte
	return builtin.lc.Post(u, data)
}
