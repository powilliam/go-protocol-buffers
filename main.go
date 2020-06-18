package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	filename := "persons"
	
	p := &Person{
		Id: 1,
		Name: "William Porto",
		Email: "powilliam19@gmail.com",
		Phones: []*Person_PhoneNumber{
			{Type: Person_MOBILE, Number: "0000000"},
		},
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalln("> Failed to encode: ", err)
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("> Failed to write persons file: ", err)
	}

	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("> Error reading file:", err)
	}
	persons := &Person{}
	if err := proto.Unmarshal(in, persons); err != nil {
		log.Fatalln("> Failed to parse address book:", err)
	}

	fmt.Println(persons)
}