package main

import (
	"fmt"
	"strings"
)

// Golang Trim() [区分大小写]
// Trim函数从字符串的两边删除预定义的字符cutset。
func main() {
	fmt.Println(strings.Trim("0120 2510", "0"))
	fmt.Println(strings.Trim("abcd axyz", "a"))
	fmt.Println(strings.Trim("abcd axyz", "A"))
	fmt.Println(strings.Trim("! Abcd dcbA !", "A"))
	fmt.Println(strings.Trim("! Abcd dcbA !", "!"))
	fmt.Println(strings.Trim(" Abcd dcbA ", " "))
}
