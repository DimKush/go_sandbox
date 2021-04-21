package BidirList

import (
	"reflect"

	"github.com/DimKush/go_sandbox/tree/main/BidirList/internal/BidirNode"
	"github.com/pkg/errors"
)

type BidirList struct {
	typeData reflect.Type
	bottom   *BidirNode.BidirNode
	top      *BidirNode.BidirNode
}

func (d *BidirList) ConstructList(sli []interface{}) error {
	//if slice is empty
	if len(sli) == 0 {
		return errors.Errorf("Wrong type detection. slice is empty.")
	}
	// detect type in first time
	if d.typeData == nil {
		d.typeData = reflect.TypeOf(sli[0])
	}
	// check type
	if reflect.TypeOf(sli[0]) != d.typeData {
		errors.Errorf("Wrong type in slice : %s. Current type of List is %s", reflect.TypeOf(sli[0]), d.typeData)
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
		errors.Errorf("Wrong type in slice : %s. Current type of List is %s", reflect.TypeOf(val), d.typeData)
	}

	if d.top != nil {
		current := d.top

		for current.GetNext() != nil {
			current = current.GetNext()
		}

		current.GetNext().SetData(val)
		current.GetNext().SetNext(d.bottom)
		current.GetNext().SetPrev(current)
		d.bottom.SetPrev(current.GetNext())
	} else {
		d.top = new(BidirNode.BidirNode)
		d.bottom = new(BidirNode.BidirNode)
		d.PushBack(val)
	}

	return nil
}

func (d *BidirList) List() reflect.Type {
	return d.typeData
}

func Len() {
	//var curNode =
}
