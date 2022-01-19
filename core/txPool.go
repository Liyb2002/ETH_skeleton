package core

import(
//	"fmt"
	"time"
	"bytes"
	"strconv"
)

type Tx struct{
	Sender string
	Recipient string
	Amount int

}
type TxPool struct{
	timeStamp time.Time
	txList [] *Tx
}


func CreateNewTx(sender string, recipient string, amount int)*Tx{
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
func ViewTxPool(pool *TxPool)string{
	thisTx :=""
	for _,tx := range (*pool).txList{
		thisTx += ViewTx(tx)
	}
	return thisTx
}

func ViewTx(tx *Tx)string{
	var buffer bytes.Buffer
	buffer.WriteString(" \n ")
	buffer.WriteString(" Sender is ")
	buffer.WriteString(tx.Sender)
	buffer.WriteString(" recipient is ")
	buffer.WriteString(tx.Recipient)
	buffer.WriteString(" amount is ")
	buffer.WriteString(strconv.Itoa(int(tx.Amount)))

	return buffer.String()
}