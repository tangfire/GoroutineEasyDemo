package main

import (
	"fmt"
	"strings"
)

// Golang Split() [区分大小写]
// 拆分功能将字符串拆分为一个切片。将s字符串拆分为用sep分隔的所有子字符串，并返回这些分隔符之间的子字符串的一部分。
func main() {
	strSlice := strings.Split("a,b,c", ",")
	fmt.Println(strSlice, "\n")

	strSlice = strings.Split("I love my country", " ")
	for _, v := range strSlice {
		fmt.Println(v)
	}

	strSlice = strings.Split("abacadaeaf", "a")
	fmt.Println("\n", strSlice)

	strSlice = strings.Split("abacadaeaf", "A")
	fmt.Println("\n", strSlice)

	strSlice = strings.Split("123023403450456056706780789", "0")
	fmt.Println("\n", strSlice)

	strSlice = strings.Split("123023403450456056706780789", ",")
	fmt.Println("\n", strSlice)
}
