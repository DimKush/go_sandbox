package BidirList

import (
	"fmt"
	"reflect"

	"errors"

	"github.com/DimKush/go_sandbox/tree/main/BidirList/internal/BidirNode"
)

type BidirList struct {
	typeData reflect.Type
	bottom   *BidirNode.BidirNode
	top      *BidirNode.BidirNode
}

func (d *BidirList) ConstructList(sli []interface{}) error {
	//if slice is empty
	if len(sli) == 0 {
		return errors.New("Wrong type detection. slice is empty.")
	}
	// detect type in first time
	if d.typeData == nil {
		d.typeData = reflect.TypeOf(sli[0])
	}
	// check type
	if reflect.TypeOf(sli[0]) != d.typeData {
		//"Wrong type in slice Current type of List is %s", reflect.TypeOf(sli[0]), d.typeData
		return errors.New("error")
	}

	for _, val := range sli {
		err := d.PushBack(val)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *BidirList) PushBack(val interface{}) error {
	// detect type in first time
	if d.typeData == nil {
		d.typeData = reflect.TypeOf(val)
	}
	// check type
	if reflect.TypeOf(val) != d.typeData {
		//errors.Errorf("Wrong type in slice : %s. Current type of List is %s", reflect.TypeOf(val), d.typeData)
		return errors.New("error")
	}

	if d.top != nil {
		current := d.top

		for current.Next != d.bottom {
			current = current.Next
		}

		current.Next = new(BidirNode.BidirNode)
		current.Next.SetData(val)
		current.Next.Prev = current
		current.Next.Next = d.bottom
		d.bottom.Prev = current.Next

	} else {

		d.top = new(BidirNode.BidirNode)
		d.bottom = new(BidirNode.BidirNode)
		d.top.Next = d.bottom
		d.PushBack(val)
	}

	return nil
}

func (d *BidirList) PushFront(val interface{}) error {
	// detect type in first time
	if d.typeData == nil {
		d.typeData = reflect.TypeOf(val)
	}
	// check type
	if reflect.TypeOf(val) != d.typeData {
		//errors.Errorf("Wrong type in slice : %s. Current type of List is %s", reflect.TypeOf(val), d.typeData)
		return errors.New("error")
	}

	if d.bottom != nil {
		current := d.bottom

		for current.Prev != d.top {
			current = current.Prev
		}

		current.Prev = new(BidirNode.BidirNode)
		current.Prev.SetData(val)
		current.Prev.Next = current
		current.Prev.Prev = d.top
		d.top.Next = current.Prev

	} else {

		d.top = new(BidirNode.BidirNode)
		d.bottom = new(BidirNode.BidirNode)
		d.top.Next = d.bottom
		d.PushFront(val)
	}

	return nil
}

func (d *BidirList) List() reflect.Type {
	return d.typeData
}

func (d *BidirList) Len() int {
	curNode := d.top.Next
	var count int
	if curNode != nil || curNode.Next != d.bottom {
		for curNode != d.bottom {
			//fmt.Println(curNode.GetData())
			count++
			curNode = curNode.Next

		}
	} else {
		return 0
	}
	return count
}

func (d *BidirList) Remove(i *BidirNode.BidirNode) {
	cur := d.top.Next

	for cur != d.bottom {
		if i == cur {
			cur.Prev.Next = cur.Next
			cur.Next.Prev = cur.Prev
			break
		}
		cur = cur.Next
	}
}

func (d *BidirList) First() *BidirNode.BidirNode {
	return d.top.Next
}

func (d *BidirList) Last() *BidirNode.BidirNode {
	return d.bottom.Prev
}

func (d *BidirList) Show() {
	cur := d.top.Next
	for cur != d.bottom {
		fmt.Printf("%v ", cur.Value())
		cur = cur.Next
	}
	fmt.Print("\n")
}
