package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var primeNumCount int32 = 0
var max_int int32 = 100000000

// var batch_size = 1000000
var concurrency = 10

func main() {
	var wg sync.WaitGroup

	batch_size := int(max_int / int32(concurrency))
	start := 0
	for i := 1; i <= concurrency; i++ {
		wg.Add(1)
		go checkPrimeGoRoutine(i, start, start+batch_size, &wg)
		start = start + batch_size
	}

	wg.Wait()
	fmt.Println("Total No. of primes ", primeNumCount)
}

func checkPrimeGoRoutine(batchNum int, start int, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	tStart := time.Now()
	fmt.Printf("Starting the batch no. %d for range %d-%d\n", batchNum, start, end)
	for i := start; i <= end; i++ {
		checkPrime(i)
	}
	fmt.Printf("Finished the batch no. %d for range %d-%d in time %s\n", batchNum, start, end, time.Since(tStart))
}

func checkPrime(num int) {
	if num&1 == 0 {
		return
	}
	for i := 3; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return
		}
	}
	atomic.AddInt32(&primeNumCount, 1)
}
