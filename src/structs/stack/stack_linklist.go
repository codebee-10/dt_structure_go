package stack

import (
	"errors"
	"fmt"
)

type LNode struct {
	Data int
	Next *LNode
}

func (p *LNode) Size() int {
	vernier := p
	count := 0
	for vernier.Next != nil {
		count++
		vernier = vernier.Next
	}
	return count
}

func (p *LNode) IsEmpty() bool {
	if p.Next != nil {
		return false
	}
	return true
}

func (p *LNode) Push(data int) {
	newNode := LNode{data, nil}
	if p.Next != nil {
		newNode.Next = p.Next
	}
	p.Next = &newNode
}

func (p *LNode) Pop() int {
	if p.IsEmpty() {
		panic(errors.New("栈已经为空"))
	}
	ret := p.Next.Data
	p.Next = p.Next.Next
	return ret
}

func (p *LNode) Top() int {
	if p.IsEmpty() {
		panic(errors.New("栈已经为空"))
	}
	return p.Next.Data
}

func (p *LNode) PrintStack() {
	vernier := p
	for vernier.Next != nil {
		fmt.Println(vernier.Next.Data)
		vernier = vernier.Next
	}
}
