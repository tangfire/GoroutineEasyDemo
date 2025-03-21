package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Golang FieldsFunc()
// FieldsFunc函数在每次运行满足f（c）的Unicode代码点c时都将字符串s断开，并返回s的切片数组。您可以使用此功能按数字或特殊字符的每个点分割字符串。
func main() {

	x := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	strArray := strings.FieldsFunc(`I love my country – I,love?my!country
                                 I, love, my – country`, x)
	for _, v := range strArray {
		fmt.Println(v)
	}

	fmt.Println("\n*****************Split by number*******************\n")

	y := func(c rune) bool {
		return unicode.IsNumber(c)
	}
	testArray := strings.FieldsFunc(`1 I love my country.2 I love my,country.3 I-love my country.4 I love my?country`, y)
	for _, w := range testArray {
		fmt.Println(w)
	}
}
