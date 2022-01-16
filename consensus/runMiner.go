package consensus

import (
	"strconv"
	"fmt"
	"eth/core"
)

func RunMiner(bc *core.Blockchain, numberOfBlocks int, numberOfMiners int, pool *core.TxPool){
	done := make(chan bool)
	jobs := make(chan string)

	for w:=0; w<numberOfMiners; w++{
		w:= w
		go func(){
			miner := CreateMiner("miner"+ strconv.Itoa(w))

			for{
				j, more := <- jobs
				if more{
					miner.Work(j, bc, pool)
				}else{
					fmt.Println("done all jobs")
					done <- true
				}
			}
		}()
	}

	for j:=0; j<numberOfBlocks; j++{
		jobs <- strconv.Itoa(j)
	}
	close(jobs)
	<- done
}