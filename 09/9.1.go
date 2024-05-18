package main

import "fmt"

func fruitMarket(f string) {
	a := map[string]int{
		"апельсины": 5,
		"яблоки":    3,
		"сливы":     1,
		"груши":     0,
	}

	v, ok := a[f]
	if ok {
		fmt.Println(f, " - ", v)
	} else {
		fmt.Println(f, " - такого нет")
	}
}

func main() {
	fruitMarket("апельсины")
	fruitMarket("что то такое")
}
