package main

import (
	"fmt"
	"runtime"
)

// B2MB converts bytes to megabytes to make the numbers easier to read [00:05:42]
func B2MB(b uint64) uint64 {
	return b / 1000 / 1000
}

// printMemStats takes a snapshot of the current memory usage and prints it [00:03:13]
func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MB", B2MB(m.Alloc))       // Current pile of Legos
	fmt.Printf("\tTotalAlloc = %v MB", B2MB(m.TotalAlloc)) // Total Legos ever used
	fmt.Printf("\tSys = %v MB", B2MB(m.Sys))         // Total room size
	fmt.Printf("\tNumGC = %v\n", m.NumGC)            // Cleanup crew visits
}

func main() {
	// 1. Check memory before we do anything [00:01:51]
	fmt.Println("--- Mem Stats Before Allocation ---")
	printMemStats()

	// 2. Create a massive "Lego build" (10 million integers) [00:02:15]
	// The underscores make the number easier to read: 10,000,000
	s := make([]int, 10_000_000)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}

	// 3. Check memory after building the giant slice [00:02:02]
	fmt.Println("\n--- Mem Stats After Allocation ---")
	printMemStats()

	// 4. Manually call the "Cleanup Crew" (Garbage Collector) [00:10:30]
	// This clears away the memory we just used
	runtime.GC()

	// 5. Check memory one last time to see it drop [00:11:01]
	fmt.Println("\n--- Mem Stats After Garbage Collection ---")
	printMemStats()
}