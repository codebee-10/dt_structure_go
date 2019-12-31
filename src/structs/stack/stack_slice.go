package stack

import (
	"errors"
	"fmt"
)

type SliceStack struct {
	Slice     []int
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
	return p.Slice[p.stackSize - 1]
}

//弹栈
func (p *SliceStack) Pop() int{
	if p.stackSize >0 {
		p.stackSize--
		ret := p.Slice[p.stackSize]
		p.Slice = p.Slice[:p.stackSize]
		return ret
	}
	panic(errors.New("stack 已经为空"))
}

//压栈
func (p *SliceStack) Push(t int) {
	p.Slice = append(p.Slice, t)
	p.stackSize = p.stackSize + 1
}

//打印栈元素
func (p *SliceStack) PrintStack() {
	for _, v := range p.Slice {
		fmt.Println(v)
	}
}
