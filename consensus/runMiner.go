package consensus

import (
	"strconv"
	"fmt"
	"eth/core"
	"time"
	"sync"
)


type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) ranking(name string, hashPw int) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name] = hashPw
}


func RunMiner(bc *core.Blockchain, numberOfBlocks int, 
	numberOfMiners int, pool *core.TxPool, done chan bool, closeRunMiner chan bool){
		var wg sync.WaitGroup

		c:= Container{
			counters: map[string]int{"genesis":0},
		}

	for w:=0; w<numberOfMiners; w++{
		wg.Add(1)
		w:= w
		
		go func(){
			ticker := time.NewTicker(100 * time.Millisecond)
			miner := CreateMiner("miner"+ strconv.Itoa(w))

			for{
				select{
				case <- ticker.C:
					MinerWork(miner, "mining", bc, pool)
				case <- done:
					hashPw := MinerReport(miner)
					c.ranking(miner.name, hashPw)
					
					ticker.Stop()
					wg.Done()
					return 
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println("hello?")
	fmt.Println(c.counters)
	closeRunMiner <- true
}