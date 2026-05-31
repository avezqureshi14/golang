package main

import (
	"fmt"
	"net/http"
	"time"
)

func leakyWorker() {
	for {
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go func() {
		fmt.Println("pprof running at http://localhost:6060/debug/pprof/")
		http.ListenAndServe(":6060", nil)
	}()

	for {
		go leakyWorker()
		time.Sleep(10 * time.Second)
	}
}
