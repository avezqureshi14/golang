package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var sink uint64

//go:noinline
func hotCalc(x uint64) uint64 {
	x ^= 1664525 + 1013904223
	x ^= x >> 13
	x ^= 0x5bd1e995
	x ^= x << 7
	return x
}

func cpuBurn() {
	// Simulates a CPU hotspot: tight loop + short sleep
	var x uint64 = 1

	for {
		for i := 0; i < 10_000_000; i++ {
			x = hotCalc(x)
		}

		sink = x
		time.Sleep(50 * time.Millisecond)
	}
}

func leak() {
	// Simulates a memory leak: repeatedly allocates 64KB and retains it
	var s [][]byte

	for {
		b := make([]byte, 64*1024)
		s = append(s, b)

		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// Start with DEFAULTS — no tuning yet
	// runtime.SetCPUProfileRate(500)
	// runtime.MemProfileRate = 64 * 1024
	// runtime.SetBlockProfileRate(1)
	// runtime.SetMutexProfileFraction(1)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go cpuBurn()
	go leak()

	select {}
}
