package initChain

import(
	"fmt"
	"time"
	"eth/consensus"
	"eth/core"
)

func RunBlockchain(bc *core.Blockchain, pool *core.TxPool, blockPeriod int){

		//blockPeriod describes how long time a block would be
		for{
		winner := consensus.RunMiner(bc, 5, pool, blockPeriod)
		fmt.Println("finshed this block, miner is", winner)
		fmt.Println("\n")
		time.Sleep(1*time.Second)
		}
	
}