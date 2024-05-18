package main

import "fmt"

var a = map[string]string{
	"груша":    "фрукт",
	"яблоко":   "фрукт",
	"апельсин": "фрукт",
	"тыква":    "овощ",
	"огурец":   "овощ",
	"помидор":  "овощ",
}

func main() {
	checkFood("груша")
	checkFood("тыква")
	checkFood("фыв")
}

func checkFood(f string) {
	v, ok := a[f]
	if ok {
		fmt.Println("это ", v)
	} else {
		fmt.Println("что-то странное…", f)
	}
}
