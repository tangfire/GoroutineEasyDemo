package main

import (
	"fmt"
	"strings"
)

// Golang Join(）[concatenate]
// Join（）函数从切片的元素返回字符串。Join将字符串Slice的元素连接起来以创建单个字符串。分隔符字符串sep指定在结果字符串中的切片元素之间放置的内容。
func main() {
	// Slice of strings
	textString := []string{"wen", "topgoer", "com"}
	fmt.Println(strings.Join(textString, "-"))

	// Slice of strings
	textNum := []string{"1", "2", "3", "4", "5"}
	fmt.Println(strings.Join(textNum, ""))
}
