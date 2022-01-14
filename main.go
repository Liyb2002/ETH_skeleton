package main

import (
	"fmt"
	"strconv"
	"eth/core"
	"eth/consensus"
//	"sync"
)

func main(){
	chain := core.CreateNewChain()
	done := make(chan bool)
	jobs := make(chan string)

	for w:=0; w<3; w++{
		w:= w
		go func(){
			miner := consensus.CreateMiner("miner"+ strconv.Itoa(w))

			for{
				j, more := <- jobs
				if more{
					miner.Work(j, chain)
				}else{
					fmt.Println("done all jobs")
					done <- true
				}
			}
		}()
	}

	for j:=0; j<20; j++{
		jobs <- strconv.Itoa(j)
	}

	close(jobs)
	<- done

	core.ViewBlockchain(chain)
}
