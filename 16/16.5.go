package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	exit := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-exit:
					fmt.Println("сложные вычисления горутины: ", i)
					fmt.Println("stop горутина: ", i)
					return
				default:
					fmt.Println("сложные вычисления горутины: ", i)
					time.Sleep(1 * time.Second)

				}
			}
		}(i)
	}

	wg.Add(1)
	go func(ch chan bool) {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")
		for i := 0; i < 10; i++ {
			ch <- true
		}
	}(exit)

	wg.Wait()
}
