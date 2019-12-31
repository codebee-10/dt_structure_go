package linklist

import (
	"fmt"
)

//单向链表
type SLNode struct {
	Data int
	Next *SLNode
}

//双向链表
type DLNode struct {
	Data int
	Pre  *DLNode
	Next *DLNode
}

func (p *SLNode) CreateSLNode(head *SLNode, nodeNum int) {
	curNode := head
	for i := 0; i < nodeNum; i++ {
		l := SLNode{i, nil}
		curNode.Next = &l
		curNode = &l
	}
}

func (p *DLNode) CreateDLNode(head *DLNode, nodeNum int) {
	curNode := head
	for i := 0; i < nodeNum; i++ {
		l := DLNode{i, nil, nil}
		curNode.Next = &l
		l.Pre = curNode
		curNode = &l
	}
}

func (p *SLNode) PrintSLNode(head *SLNode) {
	curNode := head.Next
	for {
		if curNode == nil {
			break
		}
		fmt.Println(curNode.Data)
		curNode = curNode.Next
	}
}
