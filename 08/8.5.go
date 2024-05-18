package main

import "fmt"

var Empty struct{}

func main() {
	a := map[string]int{
		"слон": 3, "бегемот": 0, "носорог": 5, "лев": 1,
	}

	fmt.Printf("%+v\n", a)
	a["бегемот"] = 2
	fmt.Printf("%+v\n", a)
}
