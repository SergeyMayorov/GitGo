package main

import "fmt"

func main() {
	a := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	fmt.Printf("%+v", a)
}
