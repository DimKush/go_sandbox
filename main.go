package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"go_sandbox/custom"
	"go_sandbox/datafile"
)

func main() {
	println("beg")
	numbers, err := datafile.GetFloats("data.txt")

	fmt.Println(reflect.TypeOf(numbers))
	fmt.Println(len(numbers))
	fmt.Println(reflect.TypeOf(datafile.GetArr()))
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, param := range numbers {
		sum += param
	}

	fmt.Println("Averange of all values", strconv.FormatFloat(sum/float64(len(numbers)), 'f', 2, 64))
	fmt.Println("Sum from all values in file ", sum)

	mapIp, err := datafile.ReadString("table.csv")

	if err != nil {
		log.Fatal("Error read ip from file")

	}

	fmt.Println(mapIp)
	custom.Process()
}
