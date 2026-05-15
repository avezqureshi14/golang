// This is the pattern of having append in loop

package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func processRequests() {
	var logs [][]byte
	for {
		data := make([]byte, 1024*100) // 100KB per request
		logs = append(logs, data)
		if len(logs) > 1000 {
			logs = nil // reset
		}
		time.Sleep(50 * time.Millisecond) // slow it down
	}
}

func main() {
	// start pprof server
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	go processRequests()

	select {} // keep app alive
}
