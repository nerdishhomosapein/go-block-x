package main

import (
	"time"

	"github.com/nerdishhomosapein/go-block-x/network"
)

// Server
// Transport Layer => TCP/UDP.
// Block
// Transactions
// key-pair

func main() {

	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	err := trLocal.Connect(trRemote)
	if err != nil {
		panic(err)
	}
	err = trRemote.Connect(trLocal)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err := trRemote.SendMessage(trLocal.Addr(), []byte("hello, world!"))
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
