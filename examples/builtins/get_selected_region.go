// Shows you how to get the currently selected region

package main

import (
	"fmt"
	"net/http/httputil"

	"github.com/coltiebaby/bastion/client/league"
	"github.com/coltiebaby/bastion/client/league/builtin"
)

func main() {
	client, err := league.CreateFromUnix()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	b := builtin.New(client)
	req := &builtin.WebSocketFormatRequest{}
	resp, err := b.WebSocketFormat(req)

	requestDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(requestDump))
}
