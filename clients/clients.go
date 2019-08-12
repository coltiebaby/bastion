// Package clients contains interfaces that tie in directly with your league client
package clients

import (
	"bytes"
	"net/http"
)

type Client interface {
	NewRequest(string, string, []byte) (*http.Request, error)
	Get(string) (*http.Response, error)
	Post(string, []byte) (*http.Response, error)
}

// Basic way to create a request to most APIs
func DefaultNewRequest(req_type, url string, data []byte) (req *http.Request, err error) {

	if data != nil {
		req, err = http.NewRequest(req_type, url, bytes.NewReader(data))
		req.Header.Add("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(req_type, url, nil)
	}

	if err != nil {
		return &http.Request{}, err
	}

	return req, err

}

type Context struct {
	Components map[string]Component `json:"contextData"`
}

func NewContext() *Context {
	return &Context{
		Components: make(map[string]Component),
	}
}

func (c *Context) AddComponent(key string, component Component) {
	c.Components[key] = component
}

type Component struct {
	Type string `json:"componentType"`
}

func NewComponent(t string) Component {
	return Component{
		Type: t,
	}
}
