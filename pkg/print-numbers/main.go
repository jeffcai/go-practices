package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	//	m := &sync.Mutex{}
	//	c := sync.NewCond(m)

	go func() {
		for i := 1; i < 100; i += 2 {
			fmt.Printf("odd number: %d \n", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 2; i <= 100; i += 2 {
			fmt.Printf("even number: %d \n", i)
		}
		wg.Done()
	}()

	wg.Wait()
}
