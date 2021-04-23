package main

import (
	"testing"

	"github.com/DimKush/go_sandbox/tree/main/BidirList/internal/BidirList"
)

func TestPushBack(t *testing.T) {
	testList := BidirList.BidirList{}
	err := testList.PushBack(10)

	if err != nil {
		t.Error(err)
	}

	if testList.First().GetData() != 10 {
		t.Error("Error in TestPushBack")
	}
}

func TestPushFront(t *testing.T) {
	testList := BidirList.BidirList{}
	err := testList.PushBack(10)

	if err != nil {
		t.Error(err)
	}

	if testList.First().GetData() != 10 {
		t.Error("Error in TestPushBack")
	}
}

func TestConstruct(t *testing.T) {
	testList := BidirList.BidirList{}
	err := testList.ConstructList([]interface{}{100, 10, 20, 30, 40, 50})

	if err != nil {
		t.Error(err)
	}
	dataTst := []interface{}{100, 10, 20, 30, 40, 50}

	var pos int
	for cur := testList.First(); cur == testList.Last(); cur = cur.Next() {
		if dataTst[pos] != cur.GetData() {
			t.Error("Error in ConstructList")
		}
	}
}
