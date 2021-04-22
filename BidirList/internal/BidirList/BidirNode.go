package BidirList

import "reflect"

type BidirNode struct {
	data     interface{}
	dataType reflect.Type
	Prev     *BidirNode
	Next     *BidirNode
}

func ConstructNode(dataIn int) *BidirNode {
	return &BidirNode{data: dataIn}
}

func (d *BidirNode) SetData(val interface{}) {

	d.data = val
}

func (d *BidirNode) Value() interface{} {
	return d.data
}

// func (d *BidirNode) SetNext(nxt *BidirNode) {
// 	d.next = nxt.next
// }

// func (d *BidirNode) CreateNext(val interface{}, prev *BidirNode, next *BidirNode) {
// 	d.next = new(BidirNode)
// 	d.next.data = val
// 	d.next.prev = prev
// 	d.next.next = next
// }

// func (d *BidirNode) SetPrev(prv *BidirNode) {
// 	d.prev = prv.prev
// }

// func (d *BidirNode) GetNext() *BidirNode {
// 	return d.next
// }

// func (d *BidirNode) GetPrev() *BidirNode {
// 	return d.prev
// }
