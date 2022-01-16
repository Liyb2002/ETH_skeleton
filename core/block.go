package core

import (
//	"fmt"
)

type block struct{
	miner string
	msg string
	prev *block
	txPool *TxPool
}

func GenesisBlock()*block{
	Genesis := block{"Genesis Block", "Hello World", nil, nil}
	return &Genesis
}

func CreateNewBlock(miner string, msg string, prev *block, txPool *TxPool) *block{
	newBlock := block{miner, msg, prev, txPool}
	return &newBlock
}