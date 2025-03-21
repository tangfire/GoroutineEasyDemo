package main

import (
	"fmt"
	"strings"
)

// Golang IndexRune()
// IndexRune函数以字符串形式返回Unicode代码点r的第一个实例的索引。如果找到，则返回以0开头的特定位置。如果找不到，则返回-1。在下面的示例中，s，t和u变量类型声明为符文。
func main() {
	var s, t, u rune
	t = 'T'
	fmt.Println(strings.IndexRune("Topgoer", t))
	fmt.Println(strings.IndexRune("topgoer", t))
	fmt.Println(strings.IndexRune("opgoer", t))
	s = 1
	fmt.Println(strings.IndexRune("5221-JAPAN", s))
	u = '1'
	fmt.Println(strings.IndexRune("5221-JAPAN", u))
}
