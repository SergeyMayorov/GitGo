package main

import (
	"errors"
	"fmt"
)

func main() {

	err1 := errors.New("ошибка1")
	fmt.Println(err1)
	err2 := fmt.Errorf("ошибка2: %w", err1)
	fmt.Println(err2)
	err3 := fmt.Errorf("ошибка3: %w", err2)
	fmt.Println(err3)

	r := errors.Is(err3, err1)
	if r {
		fmt.Println("ошибка1 есть")
	} else {
		fmt.Println("ошибка1 нет")

	}
}
