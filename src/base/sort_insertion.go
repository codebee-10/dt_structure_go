// insert sort - 插入排序

package main

import "fmt"

func insertion_sort(a []int) {
	N := len(a)
	for i:=1; i<N ; i++ {
		for j:=i; j>0 && a[j]<a[j-1] ; j-- {
			tmp := a[j]
			a[j] = a[j-1]
			a[j-1] = tmp
		}
	}
}

func main() {
	a := []int{1, 90, 50, 12, 39, 101}
	insertion_sort(a)
	fmt.Println(a)
}