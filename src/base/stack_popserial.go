//根据入栈序列判断可能的出栈序列
package main

import (
	st "../structs/stack"
	"fmt"
)

func IsPopSerial(data *[]int, push *st.LNode, pop *st.LNode) bool {
	for k, v := range *data {
		push.Push(v)
		for push.Next.Data == pop.Next.Data && !pop.Next.IsEmpty() && !push.Next.IsEmpty() {
			fmt.Println("push ----", k, ":")
			push.PrintStack()
			push.Pop()

			fmt.Println("pop ----", k, ":")
			pop.PrintStack()
			pop.Pop()
		}
	}

	if push.Next.IsEmpty() && pop.Next.IsEmpty() && pop.Next.Data == push.Next.Data {
		return true
	}
	return false
}

func main() {
	stack1 := &st.LNode{0, nil}
	data := []int{1, 2, 3, 4, 5}

	stack2 := &st.LNode{0, nil}
	stack2.Push(1)
	stack2.Push(4)
	stack2.Push(5)
	stack2.Push(2)
	stack2.Push(3)
	fmt.Println("根据stack1的入栈序列判断stack2为stack1的出栈序列结果为: ", IsPopSerial(&data, stack1, stack2))

	stack3 := &st.LNode{0, nil}
	stack3.Push(2)
	stack3.Push(4)
	stack3.Push(3)
	stack3.Push(5)
	stack3.Push(1)
	fmt.Println("根据stack1的入栈序列判断stack2为stack1的出栈序列结果为: ", IsPopSerial(&data, stack1, stack3))

}
