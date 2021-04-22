package BidirList

import (
	"fmt"
	"reflect"

	"errors"
)

type BidirList struct {
	typeData reflect.Type
	bottom   *BidirNode
	top      *BidirNode
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

		for current.next != d.bottom {
			current = current.next
		}

		current.next = new(BidirNode)
		current.next.setData(val)
		current.next.prev = current
		current.next.next = d.bottom
		d.bottom.prev = current.next

	} else {

		d.top = new(BidirNode)
		d.bottom = new(BidirNode)
		d.top.next = d.bottom
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

		for current.prev != d.top {
			current = current.prev
		}

		current.prev = new(BidirNode)
		current.prev.setData(val)
		current.prev.next = current
		current.prev.prev = d.top
		d.top.next = current.prev

	} else {

		d.top = new(BidirNode)
		d.bottom = new(BidirNode)
		d.top.next = d.bottom
		d.PushFront(val)
	}

	return nil
}

func (d *BidirList) List() reflect.Type {
	return d.typeData
}

func (d *BidirList) Len() int {
	curNode := d.top.next
	var count int
	if curNode != nil || curNode.next != d.bottom {
		for curNode != d.bottom {
			//fmt.Println(curNode.GetData())
			count++
			curNode = curNode.next

		}
	} else {
		return 0
	}
	return count
}

func (d *BidirList) Remove(i *BidirNode) {
	cur := d.top.next

	for cur != d.bottom {
		if i == cur {
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
			break
		}
		cur = cur.next
	}
}

func (d *BidirList) First() *BidirNode {
	return d.top.next
}

func (d *BidirList) Last() *BidirNode {
	return d.bottom.prev
}

func (d *BidirList) Show() {
	cur := d.top.next
	for cur != d.bottom {
		fmt.Printf("%v ", cur.GetData())
		cur = cur.next
	}
	fmt.Print("\n")
}
