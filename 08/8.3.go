package main

import "fmt"

func main() {
	a := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	fmt.Printf("%+v\n", a)
	delete(a, "бегемот")
	fmt.Printf("%+v\n", a)
}
