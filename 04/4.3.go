package main

import "fmt"

func hello(v func() string) {
	res := v()
	fmt.Println(res)
}

func main() {
	var f func() string

	f = func() string {
		return fmt.Sprint("Hello Go!")
	}

	hello(f)

}
