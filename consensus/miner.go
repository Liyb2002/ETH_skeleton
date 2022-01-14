package consensus

import (
	"fmt"
	"eth/core"
)

type miner struct{
	name string
	rewards int64
}

func CreateMiner(name string)*miner{
	return &miner{name, 0}
}

func (m miner)Work(msg string, chain *core.Blockchain){
	fmt.Println("miner", m.name, "is working!")
	core.AddBlock(chain, m.name, msg)
	m.rewards += 1

}