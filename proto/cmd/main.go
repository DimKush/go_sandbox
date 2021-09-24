package main

import (
	"fmt"
	"github.com/DimKush/go_sandbox/tree/main/proto/internal/Person"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	p := new(Person.Person)
	p.Name = "Dim"
	p.Mobile = append(p.Mobile, "8-800-555-35-35")

	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatalf("Error marshaling.")
	}

	p1 := Person.Person{}
	err = proto.Unmarshal(data, &p1)
	if err != nil {
		log.Fatalf("Error unmarshall.")
	}

	fmt.Printf("%s", p1.GetName())
	fmt.Printf("\n%s", p1.Name)
}
