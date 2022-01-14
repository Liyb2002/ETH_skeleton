package main

import (
//	"fmt"
	"eth/core"
	"eth/consensus"
)

func main(){
	chain := core.CreateNewChain()

	miner := consensus.CreateMiner("miner1")
	miner.Work("hi", chain)
	miner.Work("hi23", chain)
	miner.Work("hdsfai", chain)
	core.ViewBlockchain(chain)
}
