package core

import (
	"fmt"
)

type blockchain struct{
	blockHeight int64
	genesisBlock *block
	currentBlock *block
}


func AddBlock(bc *blockchain){
	fmt.Println("This is blockchain test")
	thisblock := CreateNewBlock("yuanbo", "helloworld", (*bc).currentBlock)
	(*bc).currentBlock = thisblock
	(*bc).blockHeight = (*bc).blockHeight+1
}

func CreateNewChain()*blockchain{
	fmt.Println("New Chain Created")
	genesis := GenesisBlock()
	bc := blockchain{0, genesis, genesis}
	return &bc
}


//View Functions
func ViewBlockHeight(bc *blockchain)int64{
	return (*bc).blockHeight
}