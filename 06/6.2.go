package main

import (
	"fmt"
	"strings"
)

type contract struct {
	ID     int
	Number string
	Date   string
}

func (c *contract) Get() string {
	r := fmt.Sprintf("{ID:%d Number:%s Date:%s}\n", c.ID, c.Number, c.Date)
	r = strings.ReplaceAll(r, "\n", "\\n")
	r = strings.ReplaceAll(r, "\t", " ")
	return r
}

func main() {
	a := contract{
		ID:     1,
		Number: "#000A101\t01",
		Date:   "2024-01-31",
	}

	fmt.Println(a.Get())
}
