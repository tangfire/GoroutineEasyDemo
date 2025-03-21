package main

import (
	"fmt"
	"strings"
)

// Golang Repeat()
// Repeat函数将字符串重复指定的次数，并返回一个新字符串，该字符串由字符串s的计数副本组成。Count指定将重复字符串的次数。必须大于或等于0。
func main() {
	textString := "China"
	repString := strings.Repeat(textString, 5)
	fmt.Println(repString)

	textString = " A " // char with space on both side
	repString = strings.Repeat(textString, 5)
	fmt.Println(repString) // Repeat space also

	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println("111" + strings.Repeat("22", 2))
	fmt.Println("111" + strings.Repeat(" ", 2))
}
