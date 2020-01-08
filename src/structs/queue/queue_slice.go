//数组实现队列
package queue

import (
	"errors"
	"fmt"
)

type SliceQueue struct {
	Slice []interface{}
	front int
	rear  int
}

//判断队列是否为空
func (p *SliceQueue) IsEmpty() bool {
	return p.front == p.rear
}

//获取队列长度
func (p *SliceQueue) Size() int {
	return p.rear - p.front
}

//获取队列首元素
func (p *SliceQueue) Top() interface{} {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空"))
	}
	return p.Slice[0]
}

//新增队列元素
func (p *SliceQueue) Push(data interface{}) {
	p.Slice = append(p.Slice, data)
	p.rear++
}

//获取队列元素
func (p *SliceQueue) Pop() interface{} {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空"))
	}
	head := p.Slice[0]
	p.front++
	p.Slice = p.Slice[1:]
	return head
}

func (p *SliceQueue) PrintQueue(queue *SliceQueue) {
	for _, v := range queue.Slice {
		fmt.Println(v)
	}
}
