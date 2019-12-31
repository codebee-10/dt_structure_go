//判断链表是有存在环
package main

import (
	"fmt"
)

type LNode struct {
	Data int
	Next *LNode
}

func IsLoop(head *LNode) *LNode{
	if head == nil || head.Next == nil {
		return head
	}

	slow := head.Next
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

func FindLoopNode(head *LNode, meetNode *LNode) *LNode {
	first := head.Next
	second := meetNode
	for first != second {
		fmt.Println(first.Data, second.Data)
		first = first.Next
		second = second.Next
	}
	return first
}

func main(){
	l1 := LNode{3, nil}
	l2 := LNode{4, nil}
	l3 := LNode{5, nil}
	l4 := LNode{6, nil}
	l5 := LNode{7, nil}
	l6 := LNode{8, nil}
	l7 := LNode{9, nil}

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6
	l6.Next = &l7
	l7.Next = &l3

	head := LNode{0, &l1}

	meetNode := IsLoop(&head)

	FindLoopNode(&head, meetNode)

}