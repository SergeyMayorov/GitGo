package main

import (
	"encoding/json"
	"fmt"
)

type st struct {
	Number   int
	Landlord string
	tenat    string
}

func main() {
	a := st{
		Number:   2,
		Landlord: "Остап",
		tenat:    "Шура",
	}
	r, err := json.Marshal(a)
	if err != nil {
		fmt.Println("ошибка чтения данных")
		return
	}
	fmt.Printf("%+w", string(r))

}
