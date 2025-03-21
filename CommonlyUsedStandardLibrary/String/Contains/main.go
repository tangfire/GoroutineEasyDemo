package main

import (
	"fmt"
	"strings"
)

// Golang字符串包含功能[区分大小写]
// 您可以使用Contains（）在字符串中搜索特定的文本/字符串/字符。它返回true或false的输出。如果在字符串2中找到字符串1，则返回true。如果在字符串2中找不到字符串1，则返回false。
func main() {
	fmt.Println(strings.Contains("我是中国人", "中国"))                //true
	fmt.Println(strings.Contains("I like golang", "like"))      //true
	fmt.Println(strings.Contains("www.topgoer.com", "topgoer")) //true
	fmt.Println(strings.Contains("www.TopgoEr.com", "topgoer")) //false
	fmt.Println(strings.Contains("www.TopgoEr com", " "))       //true
}
