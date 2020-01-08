//使用hash表从数组中找出满足a+b = c+d的两个数对
package main

import "fmt"

type Pairs struct {
	first  int
	second int
}

func FindPairs(arr []int) bool {
	sumPair := map[int]*Pairs{}
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := arr[i] + arr[j]
			if _, ok := sumPair[sum]; !ok {
				sumPair[sum] = &Pairs{i, j}
			} else {
				p := sumPair[sum]
				fmt.Printf("(%v, %v), (%v, %v)", arr[p.first], arr[p.second], arr[i], arr[j])
				fmt.Println()
				return true
			}
		}
	}
	return false
}

func main() {
	arr := []int{3, 4, 7, 10, 20, 9, 8}
	FindPairs(arr)
}
