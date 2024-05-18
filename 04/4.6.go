package main

import "fmt"

func calcFib(f []int, n int) []int {
	if n == 0 {
		return f
	}

	a := len(f)

	f = append(f, f[a-1]+f[a-2])
	f = calcFib(f, n-1)
	return f
}

func main() {
	fibonachi := make([]int, 0, 23)
	fibonachi = append(fibonachi, 0)
	fibonachi = append(fibonachi, 1)
	fibonachi = calcFib(fibonachi, 23)

	fmt.Println(fibonachi)

}
