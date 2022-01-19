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
	NewTx := core.CreateNewTx(TxList[i].Sender,TxList[i].Recipient, TxList[i].Amount)
	(*pool).AddTxToPool(NewTx)
	}
	
	ioutil.WriteFile("server/output.json", []byte(""), 0644)
	return pool


}