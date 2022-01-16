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

func AddTxToPool(pool *TxPool, txInstance *Tx){
	fmt.Println("hi")
}


//View functions
func ViewTxPool(pool *TxPool){
	fmt.Println("hi")
}

func ViewTx(tx *Tx){
	fmt.Println("sender is", tx.sender, "recipient is", tx.recipient, "amount is", tx.amount)
}