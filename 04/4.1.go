package main

import "fmt"

func hf() string {
	return fmt.Sprintln("Hello Go!")
}
func main() {
	res := hf()
	fmt.Println(res)

}
