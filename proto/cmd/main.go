package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	p := new(Person)
	p.Name = "Anton"
	p.Mobile = append(p.Mobile, "8800553535")

	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	p1 := Person{}
	err = proto.Unmarshal(data, &p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", p1)
}
