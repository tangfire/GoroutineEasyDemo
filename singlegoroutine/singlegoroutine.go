package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello go")
}

func main() {
	go hello()
	fmt.Println("hello world")
	time.Sleep(1 * time.Second)
}
