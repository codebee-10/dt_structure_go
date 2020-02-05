//divide-and-conquer 利用分治法查找最大，最小值
package main

import "fmt"

func GetMaxAndMin(arr []int) (max,min int) {
	if arr == nil {
		return 0,0
	}

	len := len(arr)
	max = arr[0]
	min = arr[0]

	for i:=0;i<len-1;i=i+2 {
		if arr[i] > arr[i+1] {
			tmp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = tmp
		}
	}

	for i:=2; i<len ; i=i+2 {
		if arr[i] < min {
			min = arr[i]
		}
	}

	for i:=3; i<len; i=i+2 {
		if arr[i] > max {
			max = arr[i]
		}
	}

	if len%2 == 1 {
		if max < arr[len-1] {
			max = arr[len-1]
		}

		if min > arr[len-1] {
			min = arr[len-1]
		}
	}

	return max,min
}
func main() {
	arr := []int{7,3,19,40,4,7,1}
	max,min := GetMaxAndMin(arr)
	fmt.Println("分治法")
	fmt.Println(max, min)
}