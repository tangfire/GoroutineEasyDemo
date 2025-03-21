package main

import (
	"fmt"
	"strings"
)

// Golang SplitAfter() [区分大小写]
// SplitAfter函数将字符串分成片。在Sep的每个实例之后，SplitAfter将S切片为所有子字符串，并返回这些子字符串的切片。
func main() {
	strSlice := strings.SplitAfter("a,b,c", ",")
	fmt.Println(strSlice, "\n")

	strSlice = strings.SplitAfter("I love my country", " ")
	for _, v := range strSlice {
		fmt.Println(v)
	}

	strSlice = strings.SplitAfter("abacadaeaf", "a")
	fmt.Println("\n", strSlice)

	strSlice = strings.SplitAfter("abacadaeaf", "A")
	fmt.Println("\n", strSlice)

	strSlice = strings.SplitAfter("123023403450456056706780789", "0")
	fmt.Println("\n", strSlice)

	strSlice = strings.SplitAfter("123023403450456056706780789", ",")
	fmt.Println("\n", strSlice)
}
