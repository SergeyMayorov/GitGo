package main

import "fmt"

func main() {
	a := 1
	do(a)
}
func do(v any) {
	a := 10
	v0, ok := v.(int)
	if ok {
		a = a + v0
		fmt.Println(a)
	} else {
		fmt.Println(" невозможности приведения к int")

	}

}
