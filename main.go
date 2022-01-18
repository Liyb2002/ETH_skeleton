package main

import (
//	"fmt"
	"eth/core"
	"eth/initChain"
	"eth/server"
//	"eth/consensus"
//	"sync"
)

func main(){

	chain := core.CreateNewChain()
	go server.StartServer(chain)
	initChain.RunBlockchain(chain, 2000)

}
