package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	out := false

	for _, v1 := range arr {
		out = false
		fmt.Println("v1: ", v1)
		for _, v2 := range arr {
			if out {
				break
			}
			fmt.Println("\tv2: ", v2)
			for _, v3 := range arr {
				if out {
					break
				}
				fmt.Println("\t\tv3: ", v3)
				for i, v4 := range arr {
					fmt.Println("\t\t\tv4: ", v4)
					if i == 1 {
						out = true
						break
					}
				}
			}
		}
	}
}
