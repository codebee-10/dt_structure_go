// 找零钱问题
// 假设有1元、2元、5元、10元、20元、50元、100元、200元面额的硬币或者纸币。现在需要N元钱，有多少种零钱组合方式


package main

import (
	"fmt"
	"strconv"
)

// 动态规划
func dp(A []int, money int) int {
	dp := make([]int, money+1) 
	dp[0] = 1
	for i:=0; i<len(A); i++ {
		for j:= A[i]; j<= money; j++ {
			dp[j] = dp[j] + dp[j - A[i]]
		}
	}
	
	return dp[money]
	// fmt.Println(dp)
}

// 暴力搜索
func vsearch(A []int, index int, aim int) int{
	result := 0
	if index == len(A) {
		if aim == 0 {
			result = 1
		}else{
			result = 0
		}
	}else{
		for i:=0;i * A[index] <= aim; i++ {
			result = result + vsearch(A, index+1, aim - i*A[index])
		}
	}
	return result
}


// 记忆搜索
var B map[string]int
func msearch(A []int, index int, aim int) int {
	key := strconv.Itoa(index) + "_" + strconv.Itoa(aim)
	if _,ok := B[key]; ok {
		return B[key]
	}

	result := 0
	if index == len(A) {
		if aim == 0 {
			result = 1
		}else{
			result = 0
		}
	}else {
		for i:=0; i*A[index] <= aim; i++ {
			result = result + msearch(A, index+1, aim - i*A[index])
		}
	}
	B[key] = result
	return result
}

func main() {
	// num := 0
	A :=[]int{1,2,5,10,20,50,100,200};
	money := 2000
	//动态规划实现
	fmt.Println(dp(A, money))
	// fmt.Println(vsearch(A, 0, money))
	B = map[string]int{}
	fmt.Println(msearch(A, 0, money))
}




