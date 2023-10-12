package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	doSomething(context.Background())
}

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		go func(num int) {
			fmt.Printf("gui vao kenh %v \n", num)
			printCh <- num
		}(num)
	}

	// time.Sleep(1 * time.Second)
	cancelCtx()

	time.Sleep(10 * time.Second)

	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		case num := <-printCh:
			time.Sleep(1 * time.Second)
			fmt.Printf("doAnother: %d\n", num)

		}
	}
}
