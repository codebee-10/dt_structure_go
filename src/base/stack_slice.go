//数组实现栈
package main

import (
	"errors"
	"fmt"
)

type SliceStack struct {
	slice []int
	stackSize int
}

//判断栈是否为空
func (p *SliceStack) IsEmpty()  bool{
	return p.stackSize == 0
}

//获取栈的大小
func (p *SliceStack) Size() int{
	return p.stackSize
}

//获取栈顶元素
func (p *SliceStack) Top() int{
	if p.IsEmpty() {
		panic(errors.New("stack 已经为空"))
	}
	return p.slice[p.stackSize - 1]
}

//弹栈
func (p *SliceStack) Pop() int{
	if p.stackSize >0 {
		p.stackSize--
		ret := p.slice[p.stackSize]
		p.slice = p.slice[:p.stackSize]
		return ret
	}
	panic(errors.New("stack 已经为空"))
}

//压栈
func (p *SliceStack) Push(t int) {
	p.slice = append(p.slice, t)
	p.stackSize = p.stackSize + 1
}

func (p *SliceStack) PrintStack() {
	for _, v := range p.slice {
		fmt.Println(v)
	}
}

func main()  {
	SliceMode()
}

func SliceMode() {
	defer func() {
		if err:=recover(); err!=nil {
			fmt.Println(err)
		}
	}()

	stack := &SliceStack{slice:make([]int,0)}
	stack.Push(1)
	stack.Push(21)
	fmt.Println("栈顶元素为:", stack.Top(), "栈大小为:", stack.Size())
	stack.Pop()
	fmt.Println("弹栈成功:", stack.Size())
	fmt.Println("栈顶元素为:", stack.Top(), "栈大小为:", stack.Size())
}