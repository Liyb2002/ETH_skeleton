package initChain

import(
	"fmt"
	"time"
	"eth/consensus"
	"eth/core"
)

func RunBlockchain(bc *core.Blockchain, pool *core.TxPool){
	for{
		done := make(chan bool)
		go consensus.RunMiner(bc, 10, 5, pool, done)
		time.Sleep(3*time.Second)
		done <- true
		close(done)
		time.Sleep(1*time.Second)
		fmt.Println("hi!")
	}
}