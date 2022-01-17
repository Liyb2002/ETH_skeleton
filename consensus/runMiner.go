package consensus

import (
	"strconv"
//	"fmt"
	"eth/core"
	"time"
)

func RunMiner(bc *core.Blockchain, numberOfBlocks int, numberOfMiners int, pool *core.TxPool, 
	done chan bool){

	for w:=0; w<numberOfMiners; w++{
		w:= w
		go func(){
			ticker := time.NewTicker(100 * time.Millisecond)
			miner := CreateMiner("miner"+ strconv.Itoa(w))

			for{
				select{
				case <- ticker.C:
					MinerWork(miner, "mining", bc, pool)
				case <- done:
					MinerReport(miner)
					ticker.Stop()
					return
				}
			}
		}()

	}
}