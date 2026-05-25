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
var currentNum int32 = 2

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

func doBatch(name string, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	for {
		x := atomic.AddInt32(&currentNum,1)
		if x > int32(MAX_INT){
			break
		}
		checkPrime(int(x))
	}
	fmt.Printf("thread %s  completed in %s \n ",name , time.Since(start))
}

func Fair() {
	var wg sync.WaitGroup

	for i := 0; i < CONCURRENCY - 1; i++{
		wg.Add(1)
		go doBatch(strconv.Itoa(i),&wg)
	}
	wg.Add(1)
	go doBatch(strconv.Itoa(CONCURRENCY-1),&wg)
	wg.Wait()
}

/*
PS D:\Go Lang\go-service\concurrency> go run main.go 
thread 5  completed in 3m0.7043202s 
 thread 4  completed in 3m0.5438064s 
 thread 8  completed in 3m0.5719825s 
 thread 2  completed in 3m0.6294063s 
 thread 0  completed in 3m0.7043202s 
 thread 9  completed in 3m0.7043202s 
 thread 7  completed in 3m0.6051961s 
 thread 3  completed in 3m0.5889803s 
 thread 1  completed in 3m0.7043202s 
 thread 6  completed in 3m0.6717904s 
 
PS D:\Go Lang\go-service\concurrency> 
*/