package main

import (
	"fmt"
	"strings"
)

// Golang LastIndex() [区分大小写]
// LastIndex函数可在字符串中搜索特定的特定文本/字符/字符串。它返回字符串中最后一个实例text / character / strin的索引。如果找到，则返回以0开头的特定位置。如果找不到，则返回-1。
func main() {
	fmt.Println(strings.LastIndex("topgoer", "o")) // position j=0,a=1,p=2,a=3
	fmt.Println(strings.LastIndex("topgoer", "G"))
	fmt.Println(strings.LastIndex("Topgoer", "go"))
	fmt.Println(strings.LastIndex("TOPGOER TOPGOER", "go"))
	fmt.Println(strings.LastIndex("1234567890 1234567890", "0"))
	fmt.Println(strings.LastIndex("1234567890 1234567890", "00"))
	fmt.Println(strings.LastIndex("1234567890 1234567890", "123"))
}
