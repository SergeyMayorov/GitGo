package main

import (
	"fmt"
	"sync"
)

var startOnce sync.Once

func start() {
	fmt.Println("я выполняюсь 1 раз и не более")

}

func func1(a int, wg *sync.WaitGroup) {
	fmt.Println("горутина :", a)
	startOnce.Do(start)
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func1(i, &wg)
	}
	wg.Wait()
}
