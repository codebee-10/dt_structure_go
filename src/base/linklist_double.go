//double linked list
package main

import(
	"fmt"
)

type LNode struct{
	Data int
	Pre *LNode
	Next *LNode
}

func CreateNode(head *LNode, nodeNum int){
	curNode := head
	for i:=0;i<nodeNum;i++ {
		l := LNode{i, nil, nil}
		curNode.Next = &l
		l.Pre = curNode
		curNode = &l
	}
}


func main(){
	head := LNode{0, nil, nil}
	CreateNode(&head, 5)

	tmp := head.Next
	for i:=0;i<5;i++{
		fmt.Print(tmp.Data, tmp.Pre.Data , "\t")
		if tmp.Next == nil {
			tmp.Next = &head
		}
		fmt.Println(tmp.Next.Data)
		tmp = tmp.Next
	}

}