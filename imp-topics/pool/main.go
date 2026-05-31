package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

type Log struct {
	Timestamp int64  `json:"ts"`
	Level     string `json:"level"`
	Message   string `json:"msg"`
}

// 🔥 Pool 1: Log struct
var logPool = sync.Pool{
	New: func() interface{} {
		return new(Log)
	},
}

// 🔥 Pool 2: Buffer
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 500; i++ {

		// ✅ reuse struct
		log := logPool.Get().(*Log)
		*log = Log{
			Timestamp: time.Now().UnixNano(),
			Level:     "INFO",
			Message:   fmt.Sprintf("user_id=%d action=login status=success", i),
		}

		// ✅ reuse buffer
		buf := bufPool.Get().(*bytes.Buffer)
		buf.Reset()

		json.NewEncoder(buf).Encode(log)

		// avoid extra string allocation
		w.Write([]byte("LOG: "))
		w.Write(buf.Bytes())

		logPool.Put(log)
		bufPool.Put(buf)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
