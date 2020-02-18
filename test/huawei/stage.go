// 台阶问题
// 有N级的台阶，你一开始在底部，每次可以向上迈最多K级台阶（最少1级），问到达第N级台阶有多少种不同方式

//例如：
//N=8;K=3
// 3 3 [2, （2）(1,1)]    = 3 + 6 =9
// 3 [5, (2,2,1),(2,1,1,1),(1,1,1,1,1)] = 38
// 0 [8, (2,2,2,2), (2,2,2,1,1), (2,2,1,1,1,1),(2,1,1,1,1,1,1),(1,1,1,1,1,1,1,1)] = 34

//N=5;K=2
// 2 2 [1 (1)] = 3
// 2 [3 (1,1,1)] = 4
// 0 [5, (1,1,1,1,1)] = 1

package main

import "fmt"

// N个阶梯(递归法), K = N
func countStep(N int) int{
	step := 0
	if N <= 1  { return 1 }
	for i:=1;i<N+1;i++ { step = step + countStep(N-i) }
	return step
}

// N个阶梯(递归法), K = L
func countStep1(N int, L int) int{
	step := 0
	if N <= 1  { return 1 }
	for i:=1;i<N+1;i++ {
		if i > L {
			break
		}
		step = step + countStep1(N-i, L)
	}

	return step
}

// N个阶梯(递归法), K 是 N中非连续的值
func countStep2(N int, L map[int]int) int{
	step := 0
	if N <= 1  { return 1 }
	for i:=1;i<N+1;i++ {
		if _, ok:=L[i]; ok {
			step = step + countStep2(N-i, L)
		}
	}

	return step
}

// N个阶梯(递归法), K = 1,2
func countStep3(N int) int{
	if N <= 2 { return N }
	step := countStep3(N-1) + countStep3(N-2)
	return step
}

// N个阶梯(遍历法), K = 1,2,3
func countStep4(N int) int{
	if N <= 3 { return N }
	s1, s2, s3 := 1, 2, 4
	sum := 0
	for i := 4; i <= N; i++ {
		sum = s1 + s2 + s3
		s1 = s2
		s2 = s3
		s3 = sum
	}
	return sum
}

// N个阶梯(遍历法), K = 1,3
func countStep5(N int) int{
	s1 := 1
	sum := 0
	tmp1 := 0
	tmp2 := 0
	for i := 1; i <= N; i++ {
		if i%2==0 {
			sum = s1 + tmp2
			tmp2 = s1
		}else{
			sum = s1 + tmp1
			tmp1 = s1
		}
		s1 = sum
	}
	return sum
}

// N个阶梯(递归法), K = 1,3
func countStep6(N int) int{
	if N < 3 { return 1 }
	if N == 3 { return 2 }
	step := countStep6(N-1) + countStep6(N-3)
	return step
}


func main() {
	//普遍的遍历法性能要比递归好很多
	N := 10
	K1 := 2
	K2 := map[int]int{1:1,3:3}
	//10个台阶,最多迈10个台阶,有多少种方法
	fmt.Println(countStep(N))
	//10个台阶,最多迈2个台阶,有多少种方法
	fmt.Println(countStep1(N, K1))
	//10个台阶,最多只能迈1和3个台阶,有多少种方法
	fmt.Println(countStep2(N, K2))
	//10个台阶,最多只能迈1和3个台阶,有多少种方法 - (性能最好)
	fmt.Println(countStep5(N))
	//10个台阶,最多只能迈1和3个台阶,有多少种方法
	fmt.Println(countStep6(N))

}

