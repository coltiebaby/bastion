// Shows the pickable skins for your character

package main

import (
	"fmt"
	"net/http/httputil"

	"github.com/coltiebaby/bastion/client/league"
	"github.com/coltiebaby/bastion/client/league/champselect"
)

func main() {
	client, err := league.NewFromExisting()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	cs := champselect.New(client)
	resp, err := cs.PickableSkins()

	requestDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(requestDump))
}
