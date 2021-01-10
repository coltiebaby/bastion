// Download a replay file
package main

import (
	"fmt"

	"github.com/coltiebaby/bastion/api/replay"
	"github.com/coltiebaby/bastion/client/league"
)

func main() {
	client, err := league.CreateFromUnix()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
