package main

import "fmt"

type contacts struct {
	Addss string
	Phone string
}

type user struct {
	ID   int
	Name string
	contacts
}

type employee struct {
	ID   int
	Name string
	contacts
}

func main() {
	u := user{
		ID:   0,
		Name: "Name 00",
		contacts: contacts{
			Addss: "add 00",
			Phone: "tel 00",
		},
	}

	e := employee{
		ID:   0,
		Name: "Name 01",
		contacts: contacts{
			Addss: "add 01",
			Phone: "tel 01",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
