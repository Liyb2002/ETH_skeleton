package server

import(
	"fmt"
	"net/http"
	"log"
	"eth/core"
	"html/template"
)

type Blockchain struct{
	info *core.Blockchain
}

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
		fmt.Println("hhhhhhhhhhhhhh")
       fmt.Println("username:", r.Form["username"])
    //    fmt.Println("password:", r.Form["password"])
    }
}


func StartServer(bc *core.Blockchain){
	thisBlockchain := &Blockchain{info:bc}
	http.HandleFunc("/view/", thisBlockchain.handleViewBlockchain)
	http.HandleFunc("/transact/", editHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}