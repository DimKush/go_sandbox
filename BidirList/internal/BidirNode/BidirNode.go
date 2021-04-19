package BidirNode

type BidirNode struct {
	data int
	prev *BidirNode
	next *BidirNode
}

func ConstructNode(dataIn int) *BidirNode {
	return &BidirNode{data: dataIn}
}
