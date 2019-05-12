package main

import (
	"fmt"
	"math/rand"
)

func genInts(i int, chInts chan int) {
	x := rand.Intn(1000)
	fmt.Printf("index: %d with rand number %d\n", i, x)
	chInts <- x
}

func main() {
	chInts := make(chan int)
	for i := 0; i < 2; i++ {
		go genInts(i, chInts)
	}
	n := <-chInts
	fmt.Printf("n: %d\n", n)

	select {
	case n := <-chInts:
		fmt.Printf("n: %d\n", n)
	}
}
