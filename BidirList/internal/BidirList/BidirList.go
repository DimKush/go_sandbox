package BidirList

import (
	"github.com/DimKush/go_sandbox/tree/main/BidirList/internal/BidirNode"
)

type BidirList struct {
	bottom *BidirNode.BidirNode
	top    *BidirNode.BidirNode
}

func (d *BidirList) ConstructList(sli []int) {
	for _, val := range sli {
		d.PushBack(val)
	}
}

func (d *BidirList) PushBack(val int) {
	if d.top != nil {
		current := d.top

		for current.GetNext() != nil {
			current = current.GetNext()
		}

	}
}
