package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func getDataFromDB(ctx context.Context, ch chan string) {
	defer close(ch)

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("cancelation received, stop fetching data")
			return
		default:
			// simulate fetching data from database
			time.Sleep(1 * time.Second)
			ch <- fmt.Sprintf("data-%d", i)
		}
	}
}

func getDataFromGoogle(ctx context.Context, ch chan string) {
	defer close(ch)

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("cancelation received, stop fetching data")
			return
		default:
			// simulate fetching data from database
			time.Sleep(1 * time.Second)
			ch <- fmt.Sprintf("data-%d", i)
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	ch1 := make(chan string)
	ch2 := make(chan string)

	wg.Add(2)
	go func() {
		defer wg.Done()
		getDataFromDB(ctx, ch1)
	}()

	go func() {
		defer wg.Done()
		getDataFromGoogle(ctx, ch2)
	}()

	for i := 0; i < 5; i++ {
		select {
		case data := <-ch1:
			fmt.Println("data from database:", data)
		case data := <-ch2:
			fmt.Println("data from google:", data)
		}
	}

	cancel()
	wg.Wait()
	fmt.Println("end of program")
}
