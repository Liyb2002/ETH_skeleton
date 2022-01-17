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

	go server.StartServer()

	NewPool := core.CreateNewTxPool()
	NewTx := core.CreateNewTx("lyb", "zws", 1)
	NewPool.AddTxToPool(NewTx)

	chain := core.CreateNewChain()
	initChain.RunBlockchain(chain, NewPool, 2000)
	core.ViewBlockchain(chain)

}
