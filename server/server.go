package server

import(
	"fmt"
	"net/http"
	"log"
	"eth/core"
)

type Blockchain struct{
	info *core.Blockchain
}

func (bc *Blockchain) handleViewBlockchain(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, core.ViewBlockchain(bc.info))
}

func StartServer(bc *core.Blockchain){
	thisBlockchain := &Blockchain{info:bc}
	http.HandleFunc("/view/", thisBlockchain.handleViewBlockchain)
    log.Fatal(http.ListenAndServe(":8080", nil))
}