package main

import "fmt"

const gC = 1
const C = 1

const (
	C1 = 1
	C2 = 2
	C3 = 3
	C4 = 4
	C5 = 5
)

func main() {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(gC)

	fmt.Println("------------------------------------------------------------------")
	const locC = 1
	fmt.Println(locC)

	fmt.Println("------------------------------------------------------------------")
	fmt.Println(C)
	const C = 2
	fmt.Println(C)

	fmt.Println("------------------------------------------------------------------")
	fmt.Println(C1, C2, C3, C4, C5)
	const (
		C1 = 1
		C2 = 2
		C3 = 3
		C4 = 4
		C5 = 5
	)

	fmt.Println(C1, C2, C3, C4, C5)

	fmt.Println("------------------------------------------------------------------")
	const n int = 5
	m := 3.4 + float64(n)
	fmt.Println(m)

	fmt.Println("------------------------------------------------------------------")
	const (
		Ciota1 = 1 << iota
		Ciota2
		Ciota3
		Ciota4
		Ciota5
	)

	fmt.Println(Ciota1, Ciota2, Ciota3, Ciota4, Ciota5)

}
