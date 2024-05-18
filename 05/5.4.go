package main

import "fmt"

func main() {
	a, b, c := "a", "b", "c"
	fmt.Println(a, &a, b, &b, c, &c)
	/*различные переменные хранятся в разных адресах в памяти*/
}
