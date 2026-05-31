package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type Log struct {
	Timestamp int64  `json:"ts"`
	Level     string `json:"level"`
	Message   string `json:"msg"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 500; i++ {

		// allocation 1
		log := Log{
			Timestamp: time.Now().UnixNano(),
			Level:     "INFO",
			Message:   fmt.Sprintf("user_id=%d action=login status=success", i),
		}

		// allocation 2
		data, _ := json.Marshal(log)

		// allocation 3 (string concat internally)
		final := fmt.Sprintf("LOG: %s\n", string(data))

		w.Write([]byte(final))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
