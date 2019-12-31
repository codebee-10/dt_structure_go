//数组实现队列
package main

import (
	"errors"
	"fmt"
)

type SliceQueue struct {
	slice []int
	front int
	rear int
}

//判断队列是否为空
func (p *SliceQueue) IsEmpty() bool{
	return p.front == p.rear
}

//获取队列长度
func (p *SliceQueue) Size()  int{
	return p.rear - p.front
}

//获取队列首元素
func (p *SliceQueue) Top()  int{
	if p.IsEmpty() {
		panic(errors.New("队列已经为空"))
	}
	return p.slice[0]
}

//新增队列元素
func (p *SliceQueue) Push(data int)  {
	p.slice = append(p.slice, data)
	p.rear++
}

//获取队列元素
func (p *SliceQueue) Pop() int{
	if p.IsEmpty() {
		panic(errors.New("队列已经为空"))
	}
	head := p.slice[p.front]
	p.front++
	p.slice = p.slice[1:]
	return head
}

func (p *SliceQueue) PrintQueue(queue *SliceQueue) {
	for _,v := range queue.slice {
		fmt.Println(v)
	}
}

func main()  {
	SliceMode()
}

func SliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	sliceQueue := &SliceQueue{slice:make([]int, 0)}
	sliceQueue.Push(1)
	sliceQueue.Push(2)
	sliceQueue.Push(3)
	sliceQueue.Push(4)
	fmt.Println("队列大小为:", sliceQueue.Size(), "队列头部元素为:", sliceQueue.Top())
	sliceQueue.PrintQueue(sliceQueue)
	fmt.Print("执行出队列: ")
	sliceQueue.Pop()
	fmt.Println("队列大小为:", sliceQueue.Size(), "队列头部元素为:", sliceQueue.Top())
	sliceQueue.PrintQueue(sliceQueue)
	fmt.Print("执行出队列: ")
	sliceQueue.Pop()
	fmt.Println("队列大小为:", sliceQueue.Size(), "队列头部元素为:", sliceQueue.Top())
	sliceQueue.PrintQueue(sliceQueue)

}