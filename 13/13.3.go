package main

import (
	"encoding/xml"
	"fmt"
)

type contract struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

type contracts struct {
	Contracts []contract `xml:"contract"`
}

const xmlData = `
<?xml version="1.0" encoding="UTF-8"?>
<contracts>
    <contract>
        <number>1</number>
        <sign_date>2023-09-02</sign_date>
        <landlord>Остап</landlord>
        <tenat>Шура</tenat>
    </contract>
    <contract>
        <number>2</number>
        <sign_date>2023-09-03</sign_date>
        <landlord>Бендер</landlord>
        <tenat>Балаганов</tenat>
    </contract>
</contracts>
`

func main() {

	var data contracts
	err := xml.Unmarshal([]byte(xmlData), &data)
	if err != nil {
		fmt.Println("Ошибка при распарсивании XML:", err)
		return
	}

	for _, i := range data.Contracts {
		fmt.Printf("%+w\n", i)
		fmt.Println("-----------------------------------------------------------")
	}
}
