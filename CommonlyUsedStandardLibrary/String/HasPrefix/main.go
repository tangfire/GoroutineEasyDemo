package main

import (
	"fmt"
	"strings"
)

// Golang HasPrefix()
// HasPrefix函数检查字符串s是否以指定的字符串开头。如果字符串S以前缀字符串开头，则返回true，否则返回false。
func main() {
	fmt.Println(strings.HasPrefix("Topgoer", "Top")) //true
	fmt.Println(strings.HasPrefix("Topgoer", "top")) //false
	fmt.Println(strings.HasPrefix("Topgoer", "ccc")) //false
	fmt.Println(strings.HasPrefix("Topgoer", ""))    //true
}
