// gcd 最大公约数 - 欧几里得算法
package main 

import "fmt"

func gcd(p int, q int) int {
	if q == 0 {
		return p
	}
	r := p%q
	return gcd(q, r)
}

func main() {
	fmt.Println(gcd(34,328))
}