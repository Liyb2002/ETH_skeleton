package core

type Tx struct{
	sender string
	recipient string
	amount int

}
type TxPool struct{
	var txList [10] *Tx
}
