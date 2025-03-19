package main

import "fmt"

func main() {
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	go func() {
		int_chan <- 1
	}()

	go func() {
		string_chan <- "hello"
	}()

	//如果多个channel同时ready，则随机选择一个执行
	select {
	case value := <-int_chan:
		fmt.Println("int:", value)
	case value := <-string_chan:
		fmt.Println("string:", value)
	}
	fmt.Println("done")

}
