package core

import (
	"fmt"
)

type blockchain struct{
	blockHeight int64

}

func Test(){
	fmt.Println("This is blockchain test")
	//blockchainTest := blockchain{0}
	thisblock := CreateNewBlock("yuanbo", "helloworld")
	fmt.Println(thisblock.miner)
}