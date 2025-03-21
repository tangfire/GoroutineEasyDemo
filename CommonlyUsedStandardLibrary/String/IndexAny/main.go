package main

import (
	"fmt"
	"strings"
)

// Golang IndexAny()
// IndexAny函数从string [left]中的chars [right]返回任何Unicode代码点的第一个实例的索引。它仅通过匹配字符串中的特定文本即可工作。如果找到，则返回以0开头的特定位置。如果找不到，则返回-1。
func main() {
	fmt.Println(strings.IndexAny("topgoer", "www"))
	fmt.Println(strings.IndexAny("topgoer", "ggg"))
	fmt.Println(strings.IndexAny("mobile", "one"))
	fmt.Println(strings.IndexAny("123456789", "4"))
	fmt.Println(strings.IndexAny("123456789", "0"))
}
