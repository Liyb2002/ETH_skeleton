package main

import (
//	"fmt"
	"eth/core"
	"eth/initChain"
//	"eth/consensus"
//	"sync"
)

func main(){

	NewPool := core.CreateNewTxPool()
	NewTx := core.CreateNewTx("lyb", "zws", 1)
	NewPool.AddTxToPool(NewTx)

	chain := core.CreateNewChain()
	initChain.RunBlockchain(chain, NewPool, 2000)
	core.ViewBlockchain(chain)

}
