package main

import (
	"fmt"
	"strconv"
	"eth/core"
	"eth/consensus"
	"sync"
)

func main(){
	chain := core.CreateNewChain()
	var wg sync.WaitGroup

	for w:=0; w<3; w++{
		w := w
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println(strconv.Itoa(w))
			name := "miner"+ strconv.Itoa(w)
			miner := consensus.CreateMiner(name)
			miner.Work("work!", chain)
		}()
	}
	wg.Wait()

	core.ViewBlockchain(chain)
}
