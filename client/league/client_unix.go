// +build linux, darwin

package league

import (
	"github.com/coltiebaby/bastion/client"
)

func NewFromExisting() (client.Client, error) {
	return CreateFromUnix()
}
