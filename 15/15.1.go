package main

import (
	"fmt"
	"sync"
)

func func1(a int, wg *sync.WaitGroup) {
	fmt.Println("горутина :", a)
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
