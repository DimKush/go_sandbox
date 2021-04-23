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
	err = testList.PushBack(20)

	if err != nil {
		t.Error(err)
	}

	if testList.Last().GetData() != 20 {
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

func TestListLen(t *testing.T) {
	testList := BidirList.BidirList{}
	err := testList.ConstructList([]interface{}{100, 10, 20, 30, 40})

	if err != nil {
		t.Error(err)
	}

	if testList.Len() != 5 {
		t.Errorf("Error in func BidirList.Len(): %d - expected; fact - %d", 5, testList.Len())
	}
}

func TestRemove(t *testing.T) {
	testList := BidirList.BidirList{}

	err := testList.ConstructList([]interface{}{100, 10, 20, 30, 40})
	if err != nil {
		t.Error(err)
	}

	rmv := testList.First()
	rmv = rmv.Next().Next()
	testList.Remove(rmv)

	dataTst := []interface{}{100, 10, 30, 40, 50}
	var pos int

	for cur := testList.First(); cur == testList.Last(); cur = cur.Next() {
		if dataTst[pos] != cur.GetData() {
			t.Error("Error in Remove()")
		}
	}
}

func TestGetData(t *testing.T) {
	testList := BidirList.BidirList{}

	err := testList.ConstructList([]interface{}{100, 10, 20, 30, 40})
	if err != nil {
		t.Error(err)
	}

	curData := testList.First().GetData()
	if curData != 100 {
		t.Errorf("Error in func BidirList.Len(): %d - expected; fact - %d", 100, curData)
	}
}
