package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

end:
	for {
		select {
		case n := <-ch:
			fmt.Printf("Received %d from a channel \n", n)
			break end
		default:
			fmt.Print("Channel is empty \n")
			ch <- 8
		}

		time.Sleep(20 * time.Millisecond)
	}
}
