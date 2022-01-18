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


	NewPool := core.CreateNewTxPool()
	NewTx := core.CreateNewTx("lyb", "zws", 1)
	NewPool.AddTxToPool(NewTx)

	chain := core.CreateNewChain()
	go server.StartServer(chain)
	initChain.RunBlockchain(chain, NewPool, 2000)

}
