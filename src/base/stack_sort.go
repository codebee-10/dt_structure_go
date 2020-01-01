//实现栈元素排序
package main

import (
	st "../structs/stack"
	"fmt"
)

//实现将栈顶元素移到栈底
func MoveStack(stack *st.LNode) {
	if stack.IsEmpty() {
		return
	}

	head := stack.Pop()
	if !stack.IsEmpty() {
		MoveStack(stack)
		end := stack.Pop()
		if head > end {
			stack.Push(end)
			stack.Push(head)
		} else {
			stack.Push(head)
			stack.Push(end)
		}

	} else {
		stack.Push(head)
	}
}

//实现栈元素翻转
func SortStack(stack *st.LNode) {
	for !stack.Next.IsEmpty() {
		MoveStack(stack)
		stack = stack.Next
	}
}

func main() {
	stack := &st.LNode{0, nil}
	stack.Push(1)
	stack.Push(3)
	stack.Push(5)
	stack.Push(2)
	stack.Push(99)
	stack.Push(21)
	stack.Push(6)
	fmt.Println("Before Sort:")
	fmt.Println("栈大小为:", stack.Size(), "栈顶元素为:", stack.Top())
	stack.PrintStack()
	fmt.Println("After Sort")
	SortStack(stack)
	fmt.Println("栈大小为:", stack.Size(), "栈顶元素为:", stack.Top())
	stack.PrintStack()
}
