package main

import (
	"fmt"
	"reflect"
)

func main() {
	var r rune
	var b byte

	fmt.Println(r)
	fmt.Println(b)

	x := 16
	y := 3
	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %s\n\n", x/y, x%y, reflect.TypeOf(x/y))
}
