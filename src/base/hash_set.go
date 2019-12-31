//hash map
package main

import "fmt"

type HashSet struct {
	Hset map[int]bool
}

func New() *HashSet{
	return &HashSet{make(map[int]bool)}
}

func (hset *HashSet) AddItem(i int) bool {
	_, found := hset.Hset[i]
	hset.Hset[i] = true
	return !found
}

func (hset *HashSet) Exist(i int) bool {
	_, found := hset.Hset[i]
	return found
}

func (hset *HashSet) DelItem(i int) {
	delete(hset.Hset, i)
}


func main()  {
	hset := New()
	hset.AddItem(10)
	hset.AddItem(11)
	hset.AddItem(13)
	hset.AddItem(11)

	hset.Exist(11)
	fmt.Println(hset.Exist(11))
	hset.DelItem(11)
	fmt.Println(hset.Exist(11))

	for k,v := range hset.Hset {
		fmt.Println(k, v)
	}
}