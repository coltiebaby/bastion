// Package clients contains interfaces that tie in directly with your league client
package client

import (
	"bytes"
	"net/http"
	"net/url"
)

type Client interface {
	NewRequest(string, url.URL, []byte) (*http.Request, error)
	URL(uri string) (url.URL, error)
	Get(url.URL) (*http.Response, error)
	Post(url.URL, []byte) (*http.Response, error)
}

// Basic way to create a request to most APIs
func DefaultNewRequest(req_type string, u url.URL, data []byte) (req *http.Request, err error) {
	raw := u.String()

	if data != nil {
		req, err = http.NewRequest(req_type, raw, bytes.NewReader(data))
		req.Header.Add("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(req_type, raw, nil)
	}

	if err != nil {
		return &http.Request{}, err
	}

	return req, err

}
