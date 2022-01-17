package initChain

import(
	"fmt"
	"time"
	"eth/consensus"
	"eth/core"
)

func RunBlockchain(bc *core.Blockchain, pool *core.TxPool){
	
		done := make(chan bool)
		closeRunMiner := make(chan bool)
		go consensus.RunMiner(bc, 10, 5, pool, done, closeRunMiner)
		time.Sleep(2*time.Second)
		done <- true
		close(done)
		time.Sleep(1*time.Second)
		<- closeRunMiner
	
}