package initChain

import(
	"fmt"
	"time"
	"eth/consensus"
	"eth/core"
)

func RunBlockchain(bc *core.Blockchain, pool *core.TxPool){
	for{
		consensus.RunMiner(bc, 10, 5, pool)
		time.Sleep(3*time.Second)
		fmt.Println("hi!")
	}
}