package main

import (
	"fmt"
	"strings"
)

// Golang Replace() [区分大小写]
// 替换功能用字符串中的某些其他字符替换某些字符。n指定要在字符串中替换的字符数。如果n小于0，则替换次数没有限制。
func main() {

	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", 0))
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", 1))
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", 2))
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", -1))
	fmt.Println(strings.Replace("1234534232132", "1", "0", -1))
	// case-sensitive
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "AN", "anese", -1))
}
