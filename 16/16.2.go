package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("стоп горутина: ", i, ctx.Err())
					return
				default:
					fmt.Println("сложные вычисления горутины: ", i)
					time.Sleep(time.Second)
				}
			}
		}(i)
	}

	wg.Wait()
}
