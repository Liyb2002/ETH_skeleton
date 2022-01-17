package consensus

import (
	"strconv"
//	"fmt"
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


func RunMiner(bc *core.Blockchain, numberOfMiners int,  pool *core.TxPool, blockPeriod int)string{

		c:= Container{
			counters: map[string]int{"genesis":0},
		}

		done := make(chan bool)
		ticker := time.NewTicker(100 * time.Millisecond)

	for w:=0; w<numberOfMiners; w++{
		w:= w
		
		go func(){
			miner := CreateMiner("miner"+ strconv.Itoa(w))

			for{
				select{
				case <- ticker.C:
					MinerWork(miner, "mining", bc, pool)
					hashPw := MinerReport(miner)
					c.ranking(miner.name, hashPw)
				case <- done:
					return 
				}
			}
		}()
	}

	time.Sleep(time.Duration(blockPeriod) * time.Millisecond)
    ticker.Stop()
    done <- true

//	fmt.Println(c.counters)
	MaxHashPw := 0;
	Winner := ""
	for key, element := range c.counters{
		if element>MaxHashPw{
			MaxHashPw=element
			Winner = key
		}
	}

	return Winner
	
}