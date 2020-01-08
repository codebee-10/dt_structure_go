package main

import (
	bt "../structs/btree"
	que "../structs/queue"
	"fmt"
)

//打印二叉树节点数据
func PrintTreeLayer(root *bt.BTree, queue *que.SliceQueue) int {
	var maxSum, preMaxSum int

	if root != nil {
		queue.Push(root)
	}

	for !queue.IsEmpty() {
		var ldata, rdata int
		childNode := queue.Pop().(*bt.BTree)
		fmt.Print(childNode.Data, " ")
		if childNode.LeftChild != nil {
			queue.Push(childNode.LeftChild)
			ldata = childNode.LeftChild.Data.(int)
			fmt.Print("left:", ldata, " ")
		}
		if childNode.RightChild != nil {
			queue.Push(childNode.RightChild)
			rdata = childNode.RightChild.Data.(int)
			fmt.Print("right:", rdata, " ")
		}

		maxSum = childNode.Data.(int) + ldata + rdata
		fmt.Println(maxSum)

		if preMaxSum <= maxSum {
			preMaxSum = maxSum
		}
	}

	return preMaxSum
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	queue := &que.SliceQueue{Slice: []interface{}{}}
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("数组:", data)
	root := bt.ArrayToBTree(data, 0, len(data)-1)
	fmt.Println(PrintTreeLayer(root, queue))
}
