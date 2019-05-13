package main

import (
	"fmt"
	"sync"
)

var cache map[int]int
var mu sync.Mutex
var wg sync.WaitGroup

func expensiveOperation(n int) int {
	// in real code this operation would be very expensive
	return n * n
}

func getCached(n int) int {
	mu.Lock()
	v, isCached := cache[n]
	mu.Unlock()
	if isCached {
		return v
	}

	v = expensiveOperation(n)

	mu.Lock()
	// without lock may cause the error - the fatal error: concurrent map writes
	cache[n] = v
	mu.Unlock()
	return v
}

func accessCache() {
	total := 0
	for i := 0; i < 5; i++ {
		n := getCached(i)
		total += n
	}
	fmt.Printf("total: %d\n", total)
	wg.Done()
}

func main() {
	cache = make(map[int]int)

	wg.Add(1)
	go accessCache()
	wg.Add(1)
	go accessCache()
	wg.Add(1)
	go accessCache()
	wg.Add(1)
	go accessCache()
	wg.Add(1)
	go accessCache()
	wg.Add(1)
	accessCache()

	//	close(chanQuit)
	fmt.Println("main finished")
	wg.Wait()
}
