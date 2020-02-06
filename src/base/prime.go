// 判断是否是质数

package main

import "fmt"

func IsPrime(N int) bool {
	if N < 2 { 
		return false 
	}
	i:= 2
	for i*i <= N {
		if N%i == 0 {
			return false;
		}
		i++
	}
	return true;
}

func main() {
	fmt.Println(IsPrime(312231))
}