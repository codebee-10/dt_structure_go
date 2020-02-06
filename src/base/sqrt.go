// 牛顿迭代法 - 计算平方根

package main

import (
	"fmt"
	"math"
)

func sqrt(N float64) float64 {
	if N < 0 {
		return math.NaN()
	}
	err := 1e-15
	t := N
	for math.Abs(t - N/t) > err * t {
		t = (N/t + t) / 2.0
	}
	return t
}

func main() {
	fmt.Println(sqrt(2336543222222))
}