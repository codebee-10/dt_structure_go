//复制二叉树
package main

import (
	bt "../structs/btree"
	"fmt"
)

func DupTree(root *bt.BTree) *bt.BTree {
	dupTree := bt.New()
	dupTree.Data = root.Data
	if root.LeftChild != nil {
		dupTree.LeftChild = DupTree(root.LeftChild)
	}
	if root.RightChild != nil {
		dupTree.RightChild = DupTree(root.RightChild)
	}

	return dupTree
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	data := []int{-1, 3, 9, 6, -7}
	fmt.Println("数组:", data)
	root := bt.ArrayToBTree(data, 0, len(data)-1)
	bt.PrintTreeMidOrder(root)
	fmt.Println("复制后Btree :")
	copyBtree := DupTree(root)
	bt.PrintTreeMidOrder(copyBtree)
}
