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

	NewTx := core.CreateNewTx("yb", "hjh", 10)
	core.ViewTx(NewTx)

}
