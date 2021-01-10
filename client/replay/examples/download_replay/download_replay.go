// Download and Launch a replay file from your client

package main

import (
	"fmt"


	"github.com/coltiebaby/bastion/api/replay"
	leagueclient "github.com/coltiebaby/bastion/client/league"
	replayclient "github.com/coltiebaby/bastion/client/replay"
)

func main() {
    matchId := "<insert match id>"

	lc, err := leagueclient.NewFromExisting()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

    rc, err := replayclient.NewReplayClient()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	rm := replay.NewReplayManager(lc, rc, matchId)

	if err = rm.DownloadVideo(); err != nil {
		fmt.Println("Error: ", err)
		return
	}

    // Uncomment this if you want to launch a replay
	// if err = rm.LoadVideoFromClient(); err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
    // }

}
