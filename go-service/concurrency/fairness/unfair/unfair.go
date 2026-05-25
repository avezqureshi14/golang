package fairness

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var MAX_INT = 100000000
var totalPrimeNumbers int32 = 0
var CONCURRENCY = 10

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

func doBatch(name string, wg *sync.WaitGroup, nstart int, nend int) {
	defer wg.Done()
	start := time.Now()
	for i := nstart; i < nend; i++ {
		checkPrime(i)
	}
	fmt.Printf("thread %s [%d , %d) completed in %s \n ",name , nstart, nend,time.Since(start))
}

func Unfair() {
	start := time.Now()

	var wg sync.WaitGroup

	nstart := 3
	batchSize := int(float64(MAX_INT) / float64(CONCURRENCY)) 

	for i := 0; i < CONCURRENCY - 1; i++{
		wg.Add(1)
		go doBatch(strconv.Itoa(i),&wg,nstart,nstart+batchSize)
		nstart += batchSize
	}
	wg.Add(1)
	go doBatch(strconv.Itoa(CONCURRENCY-1),&wg,nstart,MAX_INT)
	wg.Wait()

	fmt.Println("checking till ", MAX_INT, " found ", totalPrimeNumbers+1, " prime numbers. took ", time.Since(start))
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