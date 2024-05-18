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

	fmt.Println("------------------------------------------------------------------------------")

	uerr2 := errors.Unwrap(err3)
	fmt.Println(uerr2)
	uerr1 := errors.Unwrap(uerr2)
	fmt.Println(uerr1)

}
