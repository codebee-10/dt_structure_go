package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(1)

	go func(){
		defer wg.Done()
		fmt.Println("ok")
	}()

	wg.Wait()
}