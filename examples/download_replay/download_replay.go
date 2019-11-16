// Download a replay file
package main

import (
	"fmt"

	"github.com/coltiebaby/bastion/api/replay"
	"github.com/coltiebaby/bastion/clients/league"
)

func main() {
	client, err := league.CreateFromUnix()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	replay := replay.NewReplay(client, "3120136499")
	err = replay.Download()
	fmt.Println(err)
}
