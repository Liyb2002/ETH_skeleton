package core

import(
	"fmt"
	"time"
)

type Tx struct{
	sender string
	recipient string
	amount int64

}
type TxPool struct{
	timeStamp time.Time
	txList [] *Tx
}


func CreateNewTx(sender string, recipient string, amount int64)*Tx{
	return &Tx{sender, recipient, amount}
}

func CreateNewTxPool()*TxPool{
	gensisTx := CreateNewTx("gensis", "gensis", 0)
	return &TxPool{timeStamp:time.Now(), txList:[] *Tx{gensisTx} }
}

func (pool *TxPool)AddTxToPool(txInstance *Tx) [] *Tx{
	//Maxium tx in a pool is 9
	if len(pool.txList)<=9{
		pool.txList = append(pool.txList, txInstance)
		return pool.txList
	}else{
		return pool.txList
	}
}


//View functions
func ViewTxPool(pool *TxPool){
	for _,tx := range (*pool).txList{
		ViewTx(tx)
	}
}

func ViewTx(tx *Tx){
	fmt.Println("sender is", tx.sender, "recipient is", tx.recipient, "amount is", tx.amount)
}