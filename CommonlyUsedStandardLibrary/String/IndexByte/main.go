package main

import (
	"fmt"
	"strings"
)

// Golang IndexByte()
// IndexByte函数返回字符串中第一个字符实例的索引。如果找到，则返回以0开头的特定位置。如果找不到，则返回-1。
func main() {
	var s, t, u byte
	t = 't'
	fmt.Println(strings.IndexByte("Topgoer", t))
	fmt.Println(strings.IndexByte("topgoer", t))
	fmt.Println(strings.IndexByte("ogoer", t))
	s = 1
	fmt.Println(strings.IndexByte("5221-topgoer", s))
	u = '1'
	fmt.Println(strings.IndexByte("5221-topgoer", u))
}
