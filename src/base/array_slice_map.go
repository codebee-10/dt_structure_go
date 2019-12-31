package main

import (
	"fmt"
	"encoding/json"
	"io"
	"os"
	"bytes"
	"log"
)

type People struct {
	name  string
	child *People
}

func main() {
	log.SetOutput(os.Stdout)
	//类似链表功能的结构体
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "Ethan",
			},
		},
	}

	content, _ := json.Marshal(*relation)
	fmt.Println(content)
	fmt.Println(*relation.child.child)

	//... 可以表示
	var num = [...]int{1:10,2:30}
	fmt.Println(num)

	array := [5]*int{0:new(int), 1:new(int)}
	a := new(int)
	*a = 10
	*array[0] = 10
	array[1] = a
	
	//函数间传递数组
	// var n1 [1e6]int
	// fmt.Println(n1)

	//切片
	source := []string{"apple", "orange", "plum", "banana", "grape"}
	slice := source[2:3:3]
	slice = append(slice, "kiwi")
	slice = append(slice, "kiwi2")

	fmt.Println(source, slice)

	//append slice
	s1 := []int{1,2}
	s2 := []int{3,4}
	fmt.Println(append(s1,s2...))

	//用户自定义类型
	type user struct {
		name string
		email string
		ext string
	}

	type admin struct{
		persion user
		level int
	}

	fred := admin {
		// persion : user{
		// 	name: "Ethan",
		// 	email: "roancsu@163.com",
		// 	ext: "123",
		// },
		persion: user{"Ethan", "roancsu@163.com", "123"},
		level : 1,
	}

	fmt.Println(fred)

	//应用类型
	s3 := []int{1,2,3}
	Func1(s3)
	fmt.Println(s3)

	//标准库
	var b bytes.Buffer
	b.Write([]byte("haody"))
	fmt.Fprintf(&b, "golang\n")
	io.Copy(os.Stdout, &b)
	//log.Fatalf("ok")
	log.Println("ok")

	for _,v := range b {
		fmt.Println(v)
	}
}


func Func1(s []int) {
	s[0] = 2
}
