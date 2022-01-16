package main

import (
//	"fmt"
	"eth/core"
	"eth/consensus"
//	"sync"
)

func main(){



	NewPool := core.CreateNewTxPool()
	NewTx := core.CreateNewTx("yb", "hjh", 1)
	NewPool.AddTxToPool(NewTx)
	core.ViewTxPool(NewPool)

	chain := core.CreateNewChain()
	consensus.RunMiner(chain, 10, 5, NewPool)
	core.ViewBlockchain(chain)

}
