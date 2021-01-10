// +build linux, darwin

package clientutil

import (
	"os/user"
	"path"
)

var (
	DEFAULT_PEMFILE string = defaultPath()
)

func defaultPath() string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	return path.Join(u.HomeDir, `Documents`, `riotgames.pem`)
}
