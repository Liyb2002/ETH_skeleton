package core

import (
	"fmt"
	"time"
	"bytes"
)

type Blockchain struct{
	blockHeight int64
	genesisBlock *block
	currentBlock *block
}


func AddBlock(bc *Blockchain, miner string, msg string, pool *TxPool){
	currentTime := time.Now()
	if currentTime.Sub((*bc).currentBlock.timeStamp) > 2*time.Second{
	fmt.Println("Add a new Block!")
	thisblock := CreateNewBlock(miner, msg, (*bc).currentBlock, pool)
	(*bc).currentBlock = thisblock
	(*bc).blockHeight = (*bc).blockHeight+1
	}
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

func ViewBlockchain(bc *Blockchain)string{
	block := (*bc).currentBlock 
	var buffer bytes.Buffer


	for block!=nil{
		buffer.WriteString("Block mined by ")
		buffer.WriteString((*block).miner)
		buffer.WriteString(" with msg ")
		buffer.WriteString((*block).msg)
		buffer.WriteString("at time ")
		buffer.WriteString((*block).timeStamp.String())
		buffer.WriteString("\n")
		buffer.WriteString("Tx is")
	
		//Genesis block has no txPool
		if block.txPool!=nil{
			blockTx := ViewTxPool((*block).txPool)
			buffer.WriteString(blockTx)
		}
		block=block.prev
		
		for i:=0; i<3; i++{
			buffer.WriteString("\n")
		}
	}
	return buffer.String()

}