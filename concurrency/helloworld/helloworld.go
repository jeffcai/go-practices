package main

import (
	"fmt"
	"time"
)

func main() {
	// create new channel of type string
	ch := make(chan string)

	// start new anonymous goroutine
	go func() {
		time.Sleep(time.Second)
		// send "Hello World" to channel
		ch <- "Hello World"
	}()
	// read from channel
	msg, ok := <-ch
	fmt.Printf("msg='%s', ok='%v'\n", msg, ok)

	go mult(1, 2, ch) // first execution, non-blocking
	// read from channel
	msg, ok = <-ch
	fmt.Printf("msg='%s', ok='%v'\n", msg, ok)

	go mult(3, 4, ch) // second execution, also non-blocking
	msg, ok = <-ch
	fmt.Printf("msg='%s', ok='%v'\n", msg, ok)
}

func mult(x, y int, ch chan string) {
	fmt.Printf("%d * %d = %d\n", x, y, x*y)
	ch <- "done"
}
