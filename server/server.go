package server

import(
	"fmt"
	"net/http"
	"log"
//	"eth/core"
)

func ViewBlockchain(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "this is blockchain")
}

func StartServer(){
	http.HandleFunc("/view/", ViewBlockchain)
    log.Fatal(http.ListenAndServe(":8080", nil))
}