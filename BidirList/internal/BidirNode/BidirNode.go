package BidirNode

type BidirNode struct {
	data int
	prev *BidirNode
	next *BidirNode
}

func ConstructNode(dataIn int) *BidirNode {
	return &BidirNode{data: dataIn}
}

func (d *BidirNode) SetData(val int) {
	d.data = val
}

func (d *BidirNode) SetNext(nxt *BidirNode) {
	d.next = nxt.next
}

func (d *BidirNode) SetPrev(prv *BidirNode) {
	d.prev = prv.prev
}
