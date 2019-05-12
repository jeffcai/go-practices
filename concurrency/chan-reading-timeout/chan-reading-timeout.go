package main

import (
	"fmt"
	"time"
)

func main() {
	chResult := make(chan int, 1)
	go func() {
		time.Sleep(1 * time.Second)
		chResult <- 5
		fmt.Println("Worker finished")
	}()

	select {
	case res := <-chResult:
		fmt.Printf("Go %d form worker \n", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timed out before worker finished \n")
	}
}
