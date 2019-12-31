//展开链表列表(扁平化)
package main

import "fmt"

type LNode struct {
	Data int
	Right *LNode
	Down *LNode
}

func (p *LNode) Insert(headRef *LNode, data int) *LNode{
	newNode := &LNode{Data:data, Down:headRef}
	headRef = newNode
	return headRef
}

func merge(a *LNode, b *LNode) *LNode {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	var result *LNode
	if a.Data < b.Data {
		result = a
		result.Down = merge(a.Down, b)
	}else{
		result = b
		result.Down = merge(a, b.Down)
	}
	return  result
}

func Flatten(root *LNode) *LNode{
	if root == nil || root.Right == nil {
		return  root
	}

	root.Right = Flatten(root.Right)
	root = merge(root, root.Right)
	return root
}

func CreateNode() *LNode{
	node := &LNode{}
	node = node.Insert(nil, 31)
	node = node.Insert(nil, 8)
	node = node.Insert(nil, 6)
	node = node.Insert(nil, 3)

	node.Right = node.Insert(node.Right, 21)
	node.Right = node.Insert(node.Right, 11)

	node.Right.Right = node.Insert(node.Right.Right, 50)
	node.Right.Right = node.Insert(node.Right.Right, 22)
	node.Right.Right = node.Insert(node.Right.Right, 15)

	node.Right.Right.Right = node.Insert(node.Right.Right.Right, 55)
	node.Right.Right.Right = node.Insert(node.Right.Right.Right, 40)
	node.Right.Right.Right = node.Insert(node.Right.Right.Right, 39)
	node.Right.Right.Right = node.Insert(node.Right.Right.Right, 3)

	return node
}

func PrintNode(info string, node *LNode) {
	fmt.Println(info)
	tmp := node
	for tmp != nil {
		fmt.Println(tmp.Data, " ")
		tmp = tmp.Down
	}
}

func main()  {
	fmt.Println("链表扁平化")
	head := CreateNode()
	head = Flatten(head)
	PrintNode("扁平化链表后:", head)
}
