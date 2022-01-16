package core

import (
//	"fmt"
	"time"
)

type block struct{
	miner string
	msg string
	timeStamp time.Time
	prev *block
	txPool *TxPool
}

func GenesisBlock()*block{
	Genesis := block{"Genesis Block", "Hello World", time.Now(), nil, nil}
	return &Genesis
}

func CreateNewBlock(miner string, msg string, prev *block, pool *TxPool) *block{
	newBlock := block{miner, msg, time.Now(),prev, pool}
	return &newBlock
}