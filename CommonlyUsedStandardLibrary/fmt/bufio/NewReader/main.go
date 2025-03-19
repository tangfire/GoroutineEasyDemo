package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容")
	text, _ := reader.ReadString('\n')
	//去除字符串 text 首尾的所有空白字符
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

func main() {
	bufioDemo()
}
