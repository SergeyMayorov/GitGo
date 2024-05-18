package main

import (
	"fmt"

	v1 "github.com/SergeyMayorov/GitGo/tree/main/10.4/hello_v1"
	v2 "github.com/SergeyMayorov/GitGo/tree/main/10.4/hello_v2"
)

func main() {
	fmt.Println(v1.Hello())
	fmt.Println(v2.Hello())
}
