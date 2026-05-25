package chunking

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var MAX_INT = 100000000
var totalPrimeNumbers int32 = 0
var CONCURRENCY = 10

func chunk[T any](data []T, chunkSize int) [][]T {
	var chunks [][]T

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}

	return chunks
}

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	atomic.AddInt32(&totalPrimeNumbers, 1)
}

func SliceChunks() {
	start := time.Now()

	var wg sync.WaitGroup

	// build data
	nums := make([]int, 0, MAX_INT-3)
	for i := 3; i < MAX_INT; i++ {
		nums = append(nums, i)
	}

	// chunk data
	chunkSize := len(nums) / CONCURRENCY
	chunks := chunk(nums, chunkSize)

	for i, c := range chunks {
		wg.Add(1)

		go func(id int, data []int) {
			defer wg.Done()

			start := time.Now()
			for _, v := range data {
				checkPrime(v)
			}

			fmt.Printf("worker %d completed in %s\n", id, time.Since(start))
		}(i, c)
	}

	wg.Wait()

	fmt.Println(
		"checking till", MAX_INT,
		"found", totalPrimeNumbers+1,
		"prime numbers. took", time.Since(start),
	)
}

/*
thread 0 [3 , 10000003) completed in 1m7.1312312s
 thread 1 [10000003 , 20000003) completed in 1m46.7596404s
 thread 2 [20000003 , 30000003) completed in 2m9.3692049s
 thread 3 [30000003 , 40000003) completed in 2m22.7120502s
 thread 4 [40000003 , 50000003) completed in 2m32.6232265s
 thread 5 [50000003 , 60000003) completed in 2m42.968535s
 thread 6 [60000003 , 70000003) completed in 2m49.3670148s
 thread 7 [70000003 , 80000003) completed in 2m57.8561193s
 thread 8 [80000003 , 90000003) completed in 2m59.3351834s
 thread 9 [90000003 , 100000000) completed in 3m2.9483574s
 checking till  100000000  found  5761455  prime numbers. took  3m2.9484176s
PS D:\Go Lang\go-service\concurrency>
*/
