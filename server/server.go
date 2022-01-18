package server

import(
	"fmt"
	"net/http"
	"log"
	"eth/core"
	"html/template"
	"strconv"
)

type Blockchain struct{
	info *core.Blockchain
}

type TxJson struct{
	sender string `json:"sender"`
	recipient string `json:"recipient"`
	amount int `json:"amount"`
}

var txReceived = []TxJson{}

func (bc *Blockchain) handleViewBlockchain(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, core.ViewBlockchain(bc.info))
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, err := template.ParseFiles("server/tx.html")
		if err!=nil{
			panic(err)
		}
        t.Execute(w, nil)
    } else {
        r.ParseForm()
		TxSender := r.Form["sender"][0]
		TxRecipient := r.Form["recipient"][0]
		TxAmount, _ := strconv.Atoi (r.Form["amount"][0])

		thisTx := TxJson{TxSender, TxRecipient, TxAmount}
		txReceived=append(txReceived, thisTx)

		//Then reload page
		t, err := template.ParseFiles("server/tx.html")
		if err!=nil{
			panic(err)
		}
        t.Execute(w, nil)

    }
}


func StartServer(bc *core.Blockchain){
	thisBlockchain := &Blockchain{info:bc}
	http.HandleFunc("/view/", thisBlockchain.handleViewBlockchain)
	http.HandleFunc("/transact/", editHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}