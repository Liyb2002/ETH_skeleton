package initChain

import(
	"fmt"
	"time"
	"eth/consensus"
	"eth/core"
	"os"
	"io/ioutil"
	"encoding/json"
)

func RunBlockchain(bc *core.Blockchain, blockPeriod int){

		//blockPeriod describes how long time a block would be
		for{
		NewPool := core.CreateNewTxPool()
		NewTx := core.CreateNewTx("lyb", "zws", 1)
		NewPool.AddTxToPool(NewTx)
		NewPool = ReadTxJson(NewPool)
		
		winner := consensus.RunMiner(bc, 5, NewPool, blockPeriod)
		core.AddBlock(bc, winner, "hello world ", NewPool)
		fmt.Println("\n")
		time.Sleep(1*time.Second)
		}
	
}

func ReadTxJson(pool *core.TxPool)*core.TxPool{
	jsonFile, err := os.Open("server/output.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()

	var TxList []core.Tx
    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &TxList)

	for i:=0; i<len(TxList); i++{
	NewTx := core.CreateNewTx(TxList[0].Sender,TxList[0].Recipient, TxList[0].Amount)
	(*pool).AddTxToPool(NewTx)
	}
	return pool


}