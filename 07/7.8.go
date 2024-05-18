package main

import (
	"fmt"
)

func delByIndex(a []int, index int) []int {
	for i := index + 1; i < len(a); i++ {
		a[i-1] = a[i]
	}
	return a[:len(a)-1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	arr = delByIndex(arr, 3)
	fmt.Println(arr)

}
