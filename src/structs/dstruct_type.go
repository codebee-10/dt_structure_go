package structs

type LNode struct {
	data int
	next *LNode
}

type SliceStack struct {
	slice []int
	sliceSize int
}

type LinklistStack struct {
	head *LNode
	end *LNode
}