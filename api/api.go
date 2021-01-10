package api

import (
	"fmt"
)

type Request struct {
	Domain  string
	Version string
	Uri     string
}

func (r Request) String() string {
	return fmt.Sprintf(`/%s/%s/%s`, r.Domain, r.Version, r.Uri)
}
