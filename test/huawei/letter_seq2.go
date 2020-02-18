//求最长连续字母序列

package main

import "fmt"

func main() {
	query := "acaccbabcccb"
	count := 0
	max := []int{0,0}
	index := 0

	for k,_ := range query {
		if k>0 && (query[k] - query[k-1])<=1{
			if count == 0 {
				index = k
			}

			count++
			if count > max[1] {
				max[0] = index //开始计数的数组下标
				max[1] = count //最大字母连续计数
			}

		}else{
			//重新开始计数
			count = 0
			index = 0
		}
	}

	fmt.Println("最长连续字母序列长度:", max[1]+1)
	fmt.Println("最长连续字母序列:", query[max[0]-1:max[0]+max[1]])
}