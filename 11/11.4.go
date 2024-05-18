package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	message string
}

func (m *myFirstError) Error() string {
	return m.message
}

func main() {
	err1 := errors.New("ошибка1")
	fmt.Println(err1)
	err2 := fmt.Errorf("ошибка2: %w", err1)
	fmt.Println(err2)
	err3 := fmt.Errorf("ошибка3: %w", err2)
	fmt.Println(err3)

	myErr := &myFirstError{message: "Моя ошибочка"}

	r := errors.As(err3, &myErr)
	if r {
		fmt.Println(myErr, " есть")
	} else {
		fmt.Println(myErr, " нет")
	}
}
