package initChain

import(
	"fmt"
	"time"
	"eth/consensus"
	"eth/core"
)

func RunBlockchain(bc *core.Blockchain, blockPeriod int){

		//blockPeriod describes how long time a block would be
		for{
		NewPool := core.CreateNewTxPool()
		NewTx := core.CreateNewTx("lyb", "zws", 1)
		NewPool.AddTxToPool(NewTx)
		
		winner := consensus.RunMiner(bc, 5, NewPool, blockPeriod)
		core.AddBlock(bc, winner, "hello world ", NewPool)
		fmt.Println("\n")
		time.Sleep(1*time.Second)
		}
	
}