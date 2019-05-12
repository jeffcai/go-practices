package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	v, isClosed := <-ch
	fmt.Printf("received %d, is channel closed: %v\n", v, isClosed)
	v, isClosed = <-ch
	// it shows the channel cosed is fals, why?
	fmt.Printf("received %d, is channel closed: %v\n", v, isClosed)
}
