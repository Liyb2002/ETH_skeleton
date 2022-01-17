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
	NewTx := core.CreateNewTx("yb", "hjh", 1)
	NewPool.AddTxToPool(NewTx)
	core.ViewTxPool(NewPool)

	chain := core.CreateNewChain()
	initChain.RunBlockchain(chain, NewPool)
	core.ViewBlockchain(chain)

}
