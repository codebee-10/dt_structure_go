//从无序重复链表中删除重复项
package main

import "fmt"

type LNode struct {
	Data int
	Next *LNode
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

func DeleteNodeRepeat(head *LNode) {
	curNode := head.Next
	for curNode != nil {
		subCurNode := curNode.Next
		preNode := curNode
		for subCurNode != nil {
			if curNode.Data == subCurNode.Data {
				preNode.Next = subCurNode.Next
			}else{
				preNode = subCurNode
			}
			subCurNode = subCurNode.Next
		}
		curNode = curNode.Next
	}
}

func main() {
	l1 := LNode{3, nil}
	l2 := LNode{3, nil}
	l3 := LNode{3, nil}
	l4 := LNode{10, nil}
	l5 := LNode{3, nil}

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	head := LNode{0, &l1}
	PrintNode(&head)
	DeleteNodeRepeat(&head)
	fmt.Println("Delete Repeat item")
	PrintNode(&head)
}