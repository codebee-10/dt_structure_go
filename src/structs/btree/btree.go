//二叉树
package btree

import "fmt"

type BTree struct {
	Data       interface{}
	LeftChild  *BTree
	RightChild *BTree
}

func New() *BTree {
	return &BTree{}
}

//将数组放到二叉树中
func ArrayToBTree(arr []int, start, end int) *BTree {
	var root *BTree
	if end >= start {
		root = New()
		mid := (start + end + 1) / 2
		root.Data = arr[mid]
		root.LeftChild = ArrayToBTree(arr, start, mid-1)
		root.RightChild = ArrayToBTree(arr, mid+1, end)
	}
	return root
}

//用中序遍历的方式打印出二叉树结点的内容
func PrintTreeMidOrder(root *BTree) {
	if root == nil {
		return
	}
	//遍历root结点的左子树
	if root.LeftChild != nil {
		PrintTreeMidOrder(root.LeftChild)
	}
	//遍历root结点
	fmt.Print(root.Data, " ")
	//遍历root结点的右子树
	if root.RightChild != nil {
		PrintTreeMidOrder(root.RightChild)
	}
	//fmt.Print(root.Data, " ")
}
