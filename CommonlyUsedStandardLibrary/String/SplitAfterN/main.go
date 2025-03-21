package main

import (
	"fmt"
	"strings"
)

// Golang SplitAfterN()[区分大小写]
// SplitAfterN函数将字符串分成片。SplitAfterN在sep的每个实例之后将String切片为子字符串，并返回这些子字符串的切片。n确定要返回的子字符串数。
// n小于0：最多n个子字符串；最后一个子字符串将是未拆分的余数。
// n等于0：结果为nil（零子字符串）
// n大于0：所有子字符串
func main() {
	strSlice := strings.SplitAfterN("a,b,c", ",", 0)
	fmt.Println(strSlice, "\n")

	strSlice = strings.SplitAfterN("a,b,c", ",", 1)
	fmt.Println(strSlice, "\n")

	strSlice = strings.SplitAfterN("a,b,c", ",", 2)
	fmt.Println(strSlice, "\n")

	strSlice = strings.SplitAfterN("a,b,c", ",", 3)
	fmt.Println(strSlice, "\n")

	strSlice = strings.SplitAfterN("I love my country", " ", -1)
	for _, v := range strSlice {
		fmt.Println(v)
	}
	strSlice = strings.SplitAfterN("123023403450456056706780789", "0", 5)
	fmt.Println("\n", strSlice)
}
