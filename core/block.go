package core

import (
	"fmt"
)

type block struct{
	miner string
	msg string
	prev *block
}

func GenesisBlock()*block{
	Genesis := block{"Genesis Block", "Hello World", nil}
	return &Genesis
}

func CreateNewBlock(miner string, msg string, prev *block) *block{
	newBlock := block{miner, msg, prev}
	fmt.Println("new block!")
	return &newBlock
}