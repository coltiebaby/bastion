package api

import (
	"encoding/json"
	"http/net"
)

func UnwrapResponse(resp *http.Response, v interface{}) (err error) {
	err = json.NewDecoder(resp.Body).Decode(v)
	return err
}
