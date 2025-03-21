package main

import (
	"fmt"
	"strings"
)

//Golang Count() [区分大小写]
//此函数计算字符串中字符/字符串/文本的不重叠实例的数量。

func main() {
	fmt.Println(strings.Count("topgoer", "t"))             //1
	fmt.Println(strings.Count("Topgoer", "T"))             //1
	fmt.Println(strings.Count("Topgoer", "M"))             //0
	fmt.Println(strings.Count("Topgoer", "goer"))          // 1
	fmt.Println(strings.Count("Topgoer", "wwwTopgoercom")) // 0
	fmt.Println(strings.Count("Shell-25152", "-25"))       //1
	fmt.Println(strings.Count("Shell-25152", "-21"))       //0
	fmt.Println(strings.Count("test", ""))                 // length of string + 1   5
	fmt.Println(strings.Count("test", " "))                //0
}
