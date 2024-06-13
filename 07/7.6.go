package main

import "fmt"

func main() {
	arr := make([]int, 0, 10)
	arr = append(arr, 4, 1, 8, 9)
	fmt.Println(arr)
}
