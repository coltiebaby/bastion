// Shows you how to get the currently selected region

package main

import (
	"fmt"
	"net/http/httputil"

	"github.com/coltiebaby/bastion/clients/league"
)

func main() {
	client, err := league.CreateFromUnix()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	resp, err := client.Get(`/riotclient/get_region_locale`)

	requestDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(requestDump))
}
