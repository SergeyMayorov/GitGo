package main

import "fmt"

func main() {

	var p string = "asd"
	var s *string
	s = &p
	fmt.Println(s)
	fmt.Println(*s)

	*s = "asd1"
	fmt.Println(s)
	fmt.Println(*s)
	fmt.Println(p)

}
