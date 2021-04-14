package main

import (
	"fmt"

	"github.com/DimKush/go_sandbox/tree/main/symbolCounter/internal/counter"
)

func main() {
	str, err := counter.CountLetters("aaaabccddddde")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}

	str, err = counter.CountLetters("abcd")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}

	str, err = counter.CountLetters("45")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}

	str, err = counter.UnpackString("a4bc2d5e")
	//str, err = counter.UnpackString(`qwe\4\5`)
}
