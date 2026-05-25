package main

import (
	"fmt"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey dude")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getHello(w, r)
		return
	}
	http.Error(w, "not allowed", 405)
}

func main() {
	// so we have two choices over here http.HandleFunc and using mux.HandleFunc , the first one registers our route to a hidden global router built into Go the second one registers router strictly to ur own local mux instance
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Server is running on port 8800")
	server.ListenAndServe()
}
