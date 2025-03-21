package main

import (
	"fmt"
	"strings"
)

// Golang TrimPrefi() [区分大小写]
// TrimPrefix函数从S字符串的开头删除前缀字符串。如果S不以前缀开头，则S将原封不动地返回。
func main() {
	var s string
	s = "I love my country"
	s = strings.TrimPrefix(s, "I")
	s = strings.TrimSpace(s)
	fmt.Println(s)

	s = "I love my country"
	s = strings.TrimPrefix(s, "i")
	fmt.Println(s)

	s = "\nI-love-my-country"
	s = strings.TrimPrefix(s, "\n")
	fmt.Println(s)

	s = "\tI-love-my-country"
	s = strings.TrimPrefix(s, "\t")
	fmt.Println(s)
}
