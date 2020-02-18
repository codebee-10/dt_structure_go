// 大数乘法
// 求两个不超过200位的非负整数的积

package main

import (
	"fmt"
	"strconv"
	)

func transaction(m map[int]int) {
	for i:=0;i<len(m);i++ {
		if m[i] >= 10 {
			m[i+1] = m[i+1] + m[i]/10
			m[i] = m[i]%10
		}else {
			continue
		}
	}
}

func print(m map[int]int) {
	for i:=len(m)-1;i>=0;i--{
		fmt.Print(m[i])
	}
}

func main() {
	n1 := "340282366920938463463374607431768211456"
	n2 := "340282366920938463463374607431768211456"
	result := map[int]int{}
	outc := 0
	inc := 0

	for i:=len(n1)-1;i>=0;i-- {
		out_num, _ := strconv.Atoi(string(n1[i]))
		for j:=len(n2)-1;j>=0;j-- {
			in_num, _ := strconv.Atoi(string(n2[j]))
			result[outc+inc] = result[outc+inc] + out_num * in_num
			inc++
		}
		inc = 0
		outc++
	}

	transaction(result)
	print(result)
}