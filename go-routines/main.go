package main

import (
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	count := runtime.GOMAXPROCS(0)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println(count)

	startTime := time.Now()
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("I am Goroutine 1")

			for j := 1; j < 10_000_000; j++ {
				x := big.NewInt(100_000)
				x = new(big.Int).Mul(x, big.NewInt(int64(100_000_000)))
				x = new(big.Int).Add(x, big.NewInt(100_000_000))
				x = new(big.Int).Mul(x, big.NewInt(100_000_000))
				x = new(big.Int).Add(x, big.NewInt(100_000_000))
				x = new(big.Int).Mul(x, big.NewInt(100_000_000))
				x = new(big.Int).Add(x, big.NewInt(100_000_000))
				x = new(big.Int).Mul(x, big.NewInt(100_000_000))
				x = new(big.Int).Add(x, big.NewInt(100_000_000))
				x = new(big.Int).Mul(x, big.NewInt(100_000_000))

			}
		}

		wg.Done()
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("I am Goroutine 2")
			for j := 1; j < 10_000_000; j++ {
				x := big.NewInt(100_000)
				x = new(big.Int).Mul(x, big.NewInt(int64(100_000_000)))
				x = new(big.Int).Add(x, big.NewInt(100_000_000))
				x = new(big.Int).Mul(x, big.NewInt(100_000_000))
				x = new(big.Int).Add(x, big.NewInt(100_000_000))
				x = new(big.Int).Mul(x, big.NewInt(100_000_000))
				x = new(big.Int).Add(x, big.NewInt(100_000_000))
				x = new(big.Int).Mul(x, big.NewInt(100_000_000))
				x = new(big.Int).Add(x, big.NewInt(100_000_000))
				x = new(big.Int).Mul(x, big.NewInt(100_000_000))

			}
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("Hàm đã chạy trong %v\n", time.Now().Sub(startTime))
}
