//链表实现栈
package main

import (
	"errors"
	"fmt"
)

type LNode struct {
	Data int
	Next *LNode
}

func (p *LNode) Size(head *LNode)  int{
	count :=0
	for head.Next != nil{
		count++
		head = head.Next
	}
	return count
}

func (p *LNode) IsEmpty(head *LNode)  bool{
	if head !=nil || head.Next !=nil {
		return false
	}
	return true
}

func (p *LNode) Push(head *LNode, data int)  {
	newNode := LNode{data, nil}
	if head.Next != nil {
		newNode.Next = head.Next
	}
	head.Next = &newNode
}

func (p *LNode) Pop(head *LNode)  int{
	if head.Next != nil && head.Next.Next !=nil {
		head.Next = head.Next.Next
		head.Next.Next = nil
		return head.Next.Data
	}
	panic(errors.New("栈已经为空"))
}

func (p *LNode) Top(head *LNode)  int{
	if p.IsEmpty(head) {
		panic(errors.New("栈已经为空"))
	}
	return head.Next.Data
}

func PrintStack(head *LNode)  {
	for head.Next != nil {
		head = head.Next
	}
}

func main()  {
	defer func() {
		if err :=recover(); err!=nil {
			fmt.Println(err)
		}
	}()

	head := LNode{0, nil}
	head.Push(&head, 1)
	head.Push(&head, 2)
	PrintStack(&head)
	fmt.Println("栈大小为: ",head.Size(&head), "栈顶元素为:", head.Top(&head))
	head.Pop(&head)
	fmt.Println("栈大小为: ",head.Size(&head), "栈顶元素为:", head.Top(&head))
	PrintStack(&head)
	head.Top(&head)
	head.Pop(&head)
	fmt.Println("栈大小为: ",head.Size(&head), "栈顶元素为:", head.Top(&head))
}