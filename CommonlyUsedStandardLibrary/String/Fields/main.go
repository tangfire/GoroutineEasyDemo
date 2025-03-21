package main

import (
	"fmt"
	"strings"
)

// Golang FieldsFunc()
// FieldsFunc函数在每次运行满足f（c）的Unicode代码点c时都将字符串s断开，并返回s的切片数组。您可以使用此功能按数字或特殊字符的每个点分割字符串。
func main() {
	testString := "I love my country"
	testArray := strings.Fields(testString)
	for _, v := range testArray {
		fmt.Println(v)
	}
}
