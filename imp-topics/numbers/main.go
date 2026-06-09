package main

import (
	"fmt"
	"sync"
	"time"
)

/*
chunksize = 20
G0
0-20

start : 0
end : 20

G1
20-40

start : 20
end : 40

G2
40-60

start : 40
end : 60

G3
60-80

start : 60
end : 80

G4
80-100

start : 80
end : 100

*/

func getChunks(goRoutineId, chunkSize int) (start, end int) {
	start = goRoutineId * chunkSize
	end = start + chunkSize
	return
}

func main() {
	num := 100
	var wg sync.WaitGroup
	workers := 5
	chunkSize := num / workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			startTime := time.Now()
			defer wg.Done()
			startRange, endRange := getChunks(id, chunkSize)
			for v := startRange; v < endRange; v++ {
				fmt.Println(v)
			}
			fmt.Printf("Wroker printed from [%d %d) in %d \n", startRange, endRange, time.Since(startTime))
		}(i)
	}
	wg.Wait()
}
