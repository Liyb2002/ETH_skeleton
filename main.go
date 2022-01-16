package main

import (
//	"fmt"
	"eth/core"
//	"eth/consensus"
//	"sync"
)

func main(){
	/*
	chain := core.CreateNewChain()
	consensus.RunMiner(chain, 10, 5)
	core.ViewBlockchain(chain)
	*/
	
	NewPool := core.CreateNewTxPool()
	for i:=0; i<100; i++{
	NewTx := core.CreateNewTx("yb", "hjh", int64(i))
	NewPool.AddTxToPool(NewTx)
	}
	core.ViewTxPool(NewPool)
}
