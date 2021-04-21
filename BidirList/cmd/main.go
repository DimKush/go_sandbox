package main

import (
	"fmt"

	"github.com/DimKush/go_sandbox/tree/main/BidirList/internal/BidirList"
)

func main() {
	customList := BidirList.BidirList{}
	err := customList.ConstructList([]interface{}{12, 23, 56})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(customList.Len())
}
