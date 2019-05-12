package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sqrtWorker(chIn chan int, chOut chan int) {
	fmt.Printf("sqrtWorker started\n")
	for i := range chIn {
		sqrt := i * i
		chOut <- sqrt
	}
	fmt.Printf("sqrtWorker finished\n")
	wg.Done()
}

func main() {
	chIn := make(chan int)
	chOut := make(chan int)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go sqrtWorker(chIn, chOut)
	}

	go func() {
		chIn <- 2
		chIn <- 4
		close(chIn)
	}()

	go func() {
		wg.Wait()
		close(chOut)
	}()

	for sqrt := range chOut {
		fmt.Printf("Got sqrt: %d\n", sqrt)
	}
}
