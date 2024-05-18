package main

import "fmt"

func main() {
	f := []string{"слон", "бегемот", "осьминог"}
	a := map[string]int{
		"слон": 3, "бегемот": 0, "носорог": 5, "лев": 1,
	}
	for _, i := range f {
		v, ok := a[i]
		if ok {
			fmt.Printf("Животное : %s, количество: %d (есть в карте : %v)\n", i, v, a)
		}
	}

}
