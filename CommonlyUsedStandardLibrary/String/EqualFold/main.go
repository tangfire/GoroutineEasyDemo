package main

import (
	"fmt"
	"strings"
)

// Golang EqualFold() [不区分大小写]
// 使用EqualFold，您可以检查两个字符串是否相等。如果两个字符串相等，则返回输出true，如果两个字符串都不相等，则返回false。
func main() {
	fmt.Println(strings.EqualFold("Topgoer", "TOPGOER"))           //true
	fmt.Println(strings.EqualFold("Topgoer", "topgoer"))           //true
	fmt.Println(strings.EqualFold("Topgoer", "Topgoercom"))        //false
	fmt.Println(strings.EqualFold("Topgoer", "goer"))              //false
	fmt.Println(strings.EqualFold("Topgoer", "Topgoer & goer"))    //false
	fmt.Println(strings.EqualFold("Topgoer-1254", "topgoer-1254")) //true
	fmt.Println(strings.EqualFold(" ", " "))                       // single space both side   //true
	fmt.Println(strings.EqualFold(" ", "  "))                      // double space right side  //false
}
