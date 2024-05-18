package main

import (
	"encoding/json"
	"fmt"
)

type st struct {
	Number   int
	Landlord string
	Tenat    string
}

func main() {
	str := `{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`
	a := st{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("ошибка чтения данных")
		return
	}

	fmt.Printf("%+w", a)

}
