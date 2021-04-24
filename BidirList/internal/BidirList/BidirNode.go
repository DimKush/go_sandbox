package BidirList

type BidirNode struct {
	data interface{}
	prev *BidirNode
	next *BidirNode
}

func (d *BidirNode) setData(val interface{}) {

	d.data = val
}

func (d *BidirNode) GetData() interface{} {
	return d.data
}

func (d *BidirNode) Next() *BidirNode {
	if d.next == nil {
		return d
	}
	return d.next
}

func (d *BidirNode) Prev() *BidirNode {
	if d.prev == nil {
		return d
	}
	return d.prev
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
