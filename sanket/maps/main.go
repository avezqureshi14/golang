package main

import (
	"fmt"
	"sync"
)

var (
	globalMap = make(map[string]int)
	mu        sync.Mutex
)

var m sync.Map

func set(key string, value int) {
	// atomics are made for singe memory vriable , the work at cpu istruction level but a map is not a sinle variable internally
	mu.Lock()
	defer mu.Unlock()
	globalMap[key] = value
}

func get(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return globalMap[key]
}

func main() {
	set("count", 1)
	fmt.Println("Value of count is ", get("count"))
}
