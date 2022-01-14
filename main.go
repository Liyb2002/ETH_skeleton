package main

import (
	"fmt"
	"eth/core"
)

func main(){
	fmt.Println("Call from main")
	chain := core.CreateNewChain()
	core.AddBlock(chain)
	core.AddBlock(chain)
	core.AddBlock(chain)
	fmt.Println(core.ViewBlockHeight(chain))
}
