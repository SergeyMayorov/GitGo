package main

import (
	"fmt"

	v1 "v1"
	v2 "v2"
)

func main() {
	//https://habr.com/ru/articles/421411/
	//go get github.com/SergeyMayorov/10.4
	//go mod init mod
	//go mod tidy
	fmt.Println(v1.Hello())
	fmt.Println(v2.Hello())
}
