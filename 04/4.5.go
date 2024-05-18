package main

import "fmt"

func main() {
	defer fmt.Println("завершение работы")
	fmt.Println(helloFunc())
}

func helloFunc() string {
	return fmt.Sprint("Hello Go!")
}
