package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	sync.Mutex
	value int
}

func (c *counter) Add(x int) {
	c.Lock()
	defer c.Unlock()
	c.value += x
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("counter:", meter.value)
}
