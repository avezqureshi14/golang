package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w,"hey from net/http")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",helloHandler)
	server := &http.Server{
		Addr:":8080",
		Handler:mux ,
	}
	log.Println("Server is running on PORT 8080")
	err := server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}