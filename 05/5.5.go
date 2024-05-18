package main

import "fmt"

func main() {
	a := 5
	change(&a)
	fmt.Println(a)
}

func change(a *int) {
	*a = 8
}
