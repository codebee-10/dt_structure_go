//链表实现队列
package queue

import (
	"errors"
	"fmt"
)

type LNode struct {
	Data int
	Next *LNode
}

type LinkListQueue struct {
	head *LNode
	end  *LNode
}

//判断队列是否为空
func (p *LinkListQueue) IsEmpty() bool {
	return p.head.Next == p.end.Next
}

//获取队列长度
func (p *LinkListQueue) Size() int {
	count := 0
	vernier := p.head
	for vernier.Next != nil {
		count++
		vernier = vernier.Next
	}
	return count
}

//获取队列首元素
func (p *LinkListQueue) Top() int {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空"))
	}
	return p.head.Next.Data
}

//新增队列元素
func (p *LinkListQueue) Push(data int) {
	newNode := LNode{data, nil}
	p.end.Next = &newNode
	p.end = &newNode
}

//获取队列元素
func (p *LinkListQueue) Pop() int {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空"))
	}
	ret := p.head.Data
	p.head = p.head.Next
	return ret
}

func (p *LinkListQueue) PrintQueue() {
	vernier := p.head
	for vernier.Next != nil {
		fmt.Println(vernier.Next.Data)
		vernier = vernier.Next
	}
}
