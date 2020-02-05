//二分法查找 Binary Search

package main

import (
	"fmt"
	"sort"
)

func rank(key int, num []int) int {
	lo := 0
	hi := len(num)
	mid := (hi + lo)/2
	for lo <= hi {
		if key < num[mid] {
			hi = mid -1 
		}else if key > num[mid] {
			lo = mid + 1
		}else {
			return mid
		}
	}

	return -1
}

func main() {
	num := []int{11,2,3,5,7,22,99,0}
	sort.Ints(num)
	fmt.Println(rank(7, num))
}