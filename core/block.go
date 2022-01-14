package core

import (
	"fmt"
)

type block struct{
	miner string
	msg string
}


func CreateNewBlock(miner string, msg string) *block{
	newBlock := block{miner, msg}
	fmt.Println("new block!")
	return &newBlock
}