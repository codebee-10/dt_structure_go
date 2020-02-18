//求最长同字母序列

package main

import (
	"fmt"
	"strings"
)

func main() {

	query := "acbac"
	text := "acaccbabb"
	count := 0
	maxStr := []string{}
	for i:=0;i<len(query);i++{
		for j:=i;j<len(query);j++{
			tmp := query[j-i:j+1]
			if strings.Contains(text, tmp) {
				maxStr = append(maxStr,tmp)
				count = len(tmp)
			}
		}
	}

	fmt.Println("可匹配字母序列:",maxStr)
    fmt.Println("最长配字母序列长度为:", count)
}