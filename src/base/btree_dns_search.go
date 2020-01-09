//实现反向DNS查找缓存
package main

import "fmt"

var CharCount = 11

//Trie树

type TrieNode struct {
	IsLeaf bool
	Url    string
	Child  []*TrieNode
}

type DNSCache struct {
	root *TrieNode
}

func NewTrieNode(count int) *TrieNode {
	return &TrieNode{IsLeaf: false, Url: "", Child: make([]*TrieNode, count)}
}

func (p *DNSCache) getIndexFromRune(char rune) int {
	if char == '.' {
		return 10
	} else {
		return int(char) - '0'
	}
}

func (p *DNSCache) getRuneFromIndex(i int) rune {
	if i == 10 {
		return '.'
	} else {
		return rune('0' + i)
	}
}

func (p *DNSCache) Insert(ip, url string) {
	pCrawl := p.root
	for _, v := range []rune(ip) {
		index := p.getIndexFromRune(v)
		if pCrawl.Child[index] == nil {
			pCrawl.Child[index] = NewTrieNode(CharCount)
		}
		pCrawl = pCrawl.Child[index]
	}

	pCrawl.IsLeaf = true
	pCrawl.Url = url
}

func (p *DNSCache) SearchDNSCache(ip string) string {
	pCrawl := p.root
	for _, v := range []rune(ip) {
		index := p.getIndexFromRune(v)
		if pCrawl.Child[index] == nil {
			return ""
		}
		pCrawl = pCrawl.Child[index]
	}

	if pCrawl != nil && pCrawl.IsLeaf {
		return pCrawl.Url
	}

	return ""
}

func NewDNSCache() *DNSCache {
	return &DNSCache{root: NewTrieNode(CharCount)}
}

func main() {
	ipAdds := []string{"10.57.11.127", "121.57.61.129", "66.125.100.103"}
	urls := []string{"www.baidu.com", "www.google.com", "www.huawei.com"}
	dnsCache := NewDNSCache()
	for i, v := range ipAdds {
		dnsCache.Insert(v, urls[i])
	}
	ip := ipAdds[0]
	result := dnsCache.SearchDNSCache(ip)
	if result != "" {
		fmt.Println(ip, "--->", result)
	} else {
		fmt.Println("no results")
	}
}
