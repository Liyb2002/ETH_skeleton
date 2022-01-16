package consensus

import (
	"fmt"
	"eth/core"
	"time"
	"math/rand"
)

type miner struct{
	name string
	rewards int64
}

func CreateMiner(name string)*miner{
	return &miner{name, 0}
}

func (m miner)Work(msg string, chain *core.Blockchain, pool *core.TxPool){
	//will produce a block ranging from 0.5-2 seconds
	time.Sleep(time.Duration(rand.Int31n(1500)+500)*time.Millisecond)
	fmt.Println("Miner", m.name, "is working!")
	core.AddBlock(chain, m.name, msg, pool)
	m.rewards += 1

}