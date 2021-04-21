package BidirNode

import "reflect"

type BidirNode struct {
	data     interface{}
	dataType reflect.Type
	prev     *BidirNode
	next     *BidirNode
}

func ConstructNode(dataIn int) *BidirNode {
	return &BidirNode{data: dataIn}
}

func (d *BidirNode) SetData(val interface{}) {

	d.data = val
}

func (d *BidirNode) SetNext(nxt *BidirNode) {
	d.next = nxt.next
}

func (d *BidirNode) SetPrev(prv *BidirNode) {
	d.prev = prv.prev
}

func (d *BidirNode) GetNext() *BidirNode {
	return d.next
}

func (d *BidirNode) GetPrev() *BidirNode {
	return d.prev
}
