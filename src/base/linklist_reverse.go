//链表逆序
package main

import "fmt"

type LNode struct {
	Data int
	Next *LNode
} 

func CreateNode(head *LNode, nodeNum int) {
	curNode := head
	for i:=0; i<nodeNum; i++ {
		l := LNode{i, nil}
		curNode.Next = &l
		curNode = &l
	}
}

func PrintNode(head *LNode) {
	curNode := head
	for {
		curNode = curNode.Next
		fmt.Println(curNode.Data)
		if curNode == nil || curNode.Next == nil {
			break
		}
	}
}

func ReverseNode(head *LNode) {
	preNode := head
	curNode := preNode.Next
	firstNode := head.Next
	for {
		nextNode := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		if nextNode == nil {
			firstNode.Next = nextNode
			break
		}
		curNode = nextNode
	}
	head.Next = curNode
}

func main() {
	head := LNode{0, nil}
	fmt.Println("Before Reverse Node:")
	CreateNode(&head, 5)
	PrintNode(&head)
	//Reverse Node
	fmt.Println("Reverse Node:")
	ReverseNode(&head)
	PrintNode(&head)
}





