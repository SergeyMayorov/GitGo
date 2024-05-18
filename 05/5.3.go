package main

import "fmt"

func main() {
	var s string = "asd"
	a := &s
	*a = "qwe"

	fmt.Println(s)
}
