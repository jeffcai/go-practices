package main

import (
	"fmt"
	"strings"
)

func main() {
	address := defangIPaddr("1.1.1.1")
	fmt.Printf("defangIPaddr: %s \n", address)
}

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
