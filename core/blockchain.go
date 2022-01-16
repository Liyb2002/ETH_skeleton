package core

import (
	"fmt"
)

type Blockchain struct{
	blockHeight int64
	genesisBlock *block
	currentBlock *block
}


func AddBlock(bc *Blockchain, miner string, msg string, txPool *TxPool){
	fmt.Println("Add a new Block!")
	thisblock := CreateNewBlock(miner, msg, (*bc).currentBlock, txPool)
	(*bc).currentBlock = thisblock
	(*bc).blockHeight = (*bc).blockHeight+1
}

func CreateNewChain()*Blockchain{
	fmt.Println("New Chain Created")
	genesis := GenesisBlock()
	bc := Blockchain{0, genesis, genesis}
	return &bc
}


//View Functions
func ViewBlockHeight(bc *Blockchain)int64{
	return (*bc).blockHeight
}

func ViewBlockchain(bc *Blockchain){
	block := (*bc).currentBlock 
	for block!=nil{
		fmt.Println("Block mined by", (*block).miner, "with msg",  (*block).msg)
		block=block.prev
	}
}