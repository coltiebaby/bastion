// Download a replay file
package main

import (
	"fmt"

	"github.com/coltiebaby/go-lcu/clients/league"
	"github.com/coltiebaby/go-lcu/clients/league/replays"
)

func main() {
	client, err := league.CreateFromUnix()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	replay := replays.NewReplay(client, "3120136499")
	err = replay.Download(client)
	fmt.Println(err)
}
