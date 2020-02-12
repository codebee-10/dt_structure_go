// select sort - 选择排序

package main

import "fmt"

func selection_sort(a []int) {
	len := len(a)
	for i:=0; i<len ; i++ {
		min := i
		for j:=i+1; j<len; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		//exchange items
		if min != i {
			tmp := a[i]
			a[i] = a[min]
			a[min] = tmp
		}
	}
}

func main() {
	a := []int{1, 90, 50, 12, 39, 101}
	selection_sort(a)
	fmt.Println(a)
}