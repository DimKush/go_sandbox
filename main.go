package main

import (
	"fmt"
	"log"
	"strconv"

	"go_sandbox/datafile"
)

func main() {
	println("beg")
	numbers, err := datafile.GetFloats("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, param := range numbers {
		sum += param
	}

	fmt.Println("Averange of all values", strconv.FormatFloat(sum/float64(len(numbers)), 'f', 2, 64))
	fmt.Println("Sum from all values in file ", sum)
}
