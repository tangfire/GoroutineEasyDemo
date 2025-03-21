package main

import (
	"fmt"
	"strings"
)

// Golang ContainsAny()[区分大小写]
// 您可以使用ContainsAny（）在字符串中搜索特定的文本/字符串/字符。它返回true或false的输出。如果在字符串中找到字符的unicode，则它返回true，否则输出将为false。您可以在下面的程序中看到ContainsAny与Contains的比较。
func main() {
	fmt.Println(strings.ContainsAny("Golang", "g"))      //true
	fmt.Println(strings.ContainsAny("Golang", "l & a"))  //true
	fmt.Println(strings.ContainsAny("GolAng", "a"))      // false
	fmt.Println(strings.ContainsAny("Golang", "G"))      //true
	fmt.Println(strings.ContainsAny("GOLANG", "GOLANG")) //true
	fmt.Println(strings.ContainsAny("GOLANG", "golang")) // false
	fmt.Println(strings.ContainsAny("Shell-12541", "1")) //true
	//  Contains vs ContainsAny
	fmt.Println(strings.ContainsAny("Shell-12541", "1-2")) // true
	fmt.Println(strings.Contains("Shell-12541", "1-2"))    // false
}
