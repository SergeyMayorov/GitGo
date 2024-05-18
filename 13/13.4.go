package main

import (
	"encoding/xml"
	"fmt"
)

type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

type contracts struct {
	List []contract `xml:"contract"`
}

func main() {
	c := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		Tenat:    "Шура Балаганов",
	}

	d := contracts{}
	d.List = append(d.List, c)
	r, err := xml.Marshal(d)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(xml.Header, string(r))
}
