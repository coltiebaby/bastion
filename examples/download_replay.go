// Download a replay file
package main

import (
	"fmt"
	"net/http/httputil"

	"github.com/coltiebaby/go-lcu/clients/league"
	"github.com/coltiebaby/go-lcu/replays"
)

func main() {
	client, err := league.CreateFromUnix()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

    replay := replays.NewReplay("3120136499")
    resp, err := replay.Download(client)
	requestDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(requestDump))
}
