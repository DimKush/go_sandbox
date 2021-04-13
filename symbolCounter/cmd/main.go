package main

import (
	"fmt"

	"github.com/DimKush/go_sandbox/tree/main/symbolCounter/internal/counter"
)

func main() {
	err, _ := counter.CountLetters("f45")
	if err != nil {
		fmt.Println(err)
	}
}
