package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// https://github.com/ngoctd314/learning/blob/master/backend/golang/chapter%206.%20go%20mistakes/10.%20The%20standard%20library.md#76-timeafter-and-memory-leaks

func main() {
	ch := make(chan Event)
	go func() {
		for i := 0; ; i++ {
			ch <- Event{}
		}
	}()
	consumerTimer(ch)
}

type Event struct{}

func consumer(ch <-chan Event) {
	for {
		printAlloc()
		select {
		case event := <-ch:
			fmt.Println("recv event: ", event)
		case <-time.After(time.Hour):
			fmt.Println("warning: no message received")
			return
		}
	}
}

func consumerContext(ch <-chan Event) {
	for {
		printAlloc()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		select {
		case event := <-ch:
			cancel()
			fmt.Println("recv event: ", event)
		case <-ctx.Done():
			cancel()
			fmt.Println("warning: no message received")
			return
		}
	}
}

func consumerTimer(ch <-chan Event) {
	for {
		timerDuration := time.Second
		timer := time.NewTimer(timerDuration)

		for {
			printAlloc()
			timer.Reset(timerDuration)
			select {
			case event := <-ch:
				fmt.Println("recv event: ", event)
			case <-timer.C:
				fmt.Println("warning: no message received")
				return
			}
		}
	}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))

}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
