package main

import (
	"fmt"
)

func main() {
	p := &Person{
		Id: 1,
		Name: "William Porto",
		Email: "powilliam19@gmail.com",
		Phones: []*Person_PhoneNumber{
			{Type: Person_MOBILE, Number: "0000000"},
		},
	}

	fmt.Println(p)
}