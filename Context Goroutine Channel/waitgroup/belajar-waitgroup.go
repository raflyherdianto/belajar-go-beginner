package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(who string) {
			defer wg.Done()

			fmt.Println("hello", who)
		}(fmt.Sprintf("john wick #%d", i))
	}
	wg.Wait()
	fmt.Println("end of program")
}
