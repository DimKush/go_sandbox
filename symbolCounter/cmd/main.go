package main

import (
	"fmt"

	"github.com/DimKush/go_sandbox/tree/main/symbolCounter/internal/counter"
)

func main() {
	str, err := counter.PackLetters("aaaabccddddde")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}

	str, err = counter.UnpackString("a4bc2d5e")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
	str, err = counter.UnpackString("abcd")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
	str, err = counter.UnpackString("45")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
	str, err = counter.UnpackString(`qwe\4\5`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
	str, err = counter.UnpackString(`qwe\45`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
	str, err = counter.UnpackString(`qwe\\5`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}

}
