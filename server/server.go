package server

import(
	"fmt"
	"net/http"
	"log"
	"eth/core"
	"html/template"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

type Blockchain struct{
	info *core.Blockchain
}

type TxJson struct{
	Sender string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount int `json:"amount"`
}

var txReceived = []TxJson{}

func (bc *Blockchain) handleViewBlockchain(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, core.ViewBlockchain(bc.info))
}

func editHandler(w http.ResponseWriter, r *http.Request) {
 //   fmt.Println("method:", r.Method) 
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

		txReceived=append(txReceived, TxJson{TxSender, TxRecipient, TxAmount})
		TxJsonW, _ := json.Marshal(txReceived)
		err := ioutil.WriteFile("server/output.json", TxJsonW, 0644)
		if err!=nil{
			panic(err)
		}

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