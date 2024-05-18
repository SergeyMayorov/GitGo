package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	close(ch)
	stop := make(chan struct{})
	go func() {
		<-ch
		stop <- struct{}{}
	}()
	<-stop
	fmt.Println("happy end")
}
