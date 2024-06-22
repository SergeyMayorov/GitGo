package main

import (
	"fmt"

	v1 "v1.0.0"
	v2 "v1.1.0"
)

func main() {
	//https://habr.com/ru/articles/421411/
	//https://go.dev/doc/modules/gomod-ref
	//go get github.com/SergeyMayorov/10.4
	//go get github.com/SergeyMayorov/10.4@v1.0.0
	//go get github.com/SergeyMayorov/10.4@v1.1.0
	//go mod init mod
	//go mod tidy
	fmt.Println(v1.Hello())
	fmt.Println(v2.Hello())
}
