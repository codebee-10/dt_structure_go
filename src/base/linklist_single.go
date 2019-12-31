//singly linked list - 单向链表
package main

import (
	"fmt"
)

type LNode struct {
	Data int
	Next *LNode
}

func CreateNode(head *LNode, nodeNum int){
	curNode := head
	for i:=0;i<nodeNum;i++ {
		l := LNode{i, nil}
		curNode.Next = &l
		curNode = &l
	}
}

func PrintNode(head *LNode) {
	curNode := head.Next
	for {
		if curNode == nil {
			break 
		}
		fmt.Println(curNode.Data)
		curNode = curNode.Next
	}
}

func main() {
	head := LNode{0, nil}
	CreateNode(&head, 5)
	PrintNode(&head)
}

