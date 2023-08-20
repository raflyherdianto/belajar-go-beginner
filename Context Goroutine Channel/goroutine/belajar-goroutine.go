package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	go fmt.Print("hello ")
	fmt.Print("world!\n")
}
