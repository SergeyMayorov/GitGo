package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr := append(arr1, arr2...)
	fmt.Println(arr)
}
