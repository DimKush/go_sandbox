package main

import (
	"fmt"
	"log"
	"reflect"
)

type X struct {
	data     interface{}
	dataType reflect.Type
}

func (x *X) showType() {
	fmt.Println(reflect.TypeOf(x.data))
}

func (x *X) insert(i interface{}) {
	if x.data == nil {
		x.dataType = reflect.TypeOf(i)
		x.data = i
	}

	if reflect.TypeOf(i) != x.dataType {
		log.Fatal("Wrong type")
	}

	x.data = i
}

func main() {
	var v X
	v.insert(10)
	v.insert("10")
	v.showType()
}
