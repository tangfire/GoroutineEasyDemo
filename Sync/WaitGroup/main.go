package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("hello world")
}

// 需要注意sync.WaitGroup是一个结构体，传递的时候要传递指针。
func main() {
	wg.Add(1)
	go hello()
	fmt.Println("hello go")
	wg.Wait()
}
