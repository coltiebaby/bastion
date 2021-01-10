// Shows you how to get the currently selected region

package main

import (
	"fmt"
	"net/http/httputil"

	"github.com/coltiebaby/bastion/client/league"
)

func main() {
	client, err := league.NewFromExisting()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	u, _ := client.URL(`/riotclient/get_region_locale`)
	resp, err := client.Get(u)

	requestDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(requestDump))
}
