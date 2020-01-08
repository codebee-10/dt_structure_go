//把数组放到二叉树中
package main

import (
	bt "../structs/btree"
	que "../structs/queue"
	"fmt"
)

//打印二叉树节点数据
func PrintTreeLayer(root *bt.BTree, queue *que.SliceQueue) {
	if root != nil {
		queue.Push(root)
	}

	for !queue.IsEmpty() {
		childNode := queue.Pop().(*bt.BTree)
		fmt.Print(childNode.Data, " ")
		if childNode.LeftChild != nil {
			queue.Push(childNode.LeftChild)
		}
		if childNode.RightChild != nil {
			queue.Push(childNode.RightChild)
		}
	}
}

//空间复杂度为O(1)打印层级Node:
func PrintTreeLayerAtO1(root *bt.BTree, level int) int {
	if root == nil || level < 0 {
		return 0
	} else if level == 0 {
		fmt.Print(root.Data, " ")
		return 1
	} else {
		fmt.Print(root.Data, " ")
		return PrintTreeLayerAtO1(root.LeftChild, level-1) + PrintTreeLayerAtO1(root.RightChild, level-1)
	}
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
	bt.PrintTreeMidOrder(root)

	fmt.Println("O(1)打印层级Node:")
	PrintTreeLayerAtO1(root, 2)

	fmt.Println("")
	PrintTreeLayer(root, queue)
	fmt.Println("")
	fmt.Println(queue.Pop())
}
