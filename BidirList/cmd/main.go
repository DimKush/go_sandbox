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
	customList.PushBack(100)
	customList.PushBack(1000)
	customList.PushFront(10)
	customList.PushFront(1000)

	cur := customList.Last()
	cur = cur.Prev()
	customList.Remove(cur)

	customList.Show()
	fmt.Println(customList.Len())
	fmt.Println(customList.List())
}
