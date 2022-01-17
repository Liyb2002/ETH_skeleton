package consensus

import (
	"fmt"
	"eth/core"
	"time"
	"math/rand"
)

type miner struct{
	name string
	hashPw int
	rewards int64
}

func CreateMiner(name string)*miner{
	return &miner{name, 0, 0}
}

func MinerWork(m *miner, msg string, chain *core.Blockchain, pool *core.TxPool){
	//will produce a block ranging from 0.5-2 seconds
	time.Sleep(time.Duration(rand.Int31n(1500)+500)*time.Millisecond)
	fmt.Println("Miner", (*m).name, "is working!")
	(*m).hashPw += rand.Intn(100)
	//core.AddBlock(chain, m.name, msg, pool)
	//m.rewards += 1

}

func MinerReport(m *miner)int{
	return (*m).hashPw
}