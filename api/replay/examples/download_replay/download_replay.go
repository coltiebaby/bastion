// Download and Launch a replay file from your client

// TODO: use the client api to pull your most recent game (too lazy)
// open client > go to your match history > click a game > click view on web
// Use the number where "99999999" is in this link

// https://matchhistory.na.leagueoflegends.com/en/#match-details/NA1/6666666666/99999999

package main

import (
	"fmt"

	"github.com/coltiebaby/bastion/api/replay"
	rc "github.com/coltiebaby/bastion/client/replay"
)

func main() {
	matchId := "99999999"

	client, err := rc.New()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	r := replay.New(client, matchId)

	if err = r.Download(); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Uncomment this if you want to launch a replay
	// if err = r.Watch(); err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

}
