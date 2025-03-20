package main

import (
	"fmt"
	"os"
)

func main() {
	var buf [16]byte
	// 读取时获取实际读取长度
	n, _ := os.Stdin.Read(buf[:])

	// 正确输出到标准输出
	os.Stdout.WriteString(string(buf[:n]))

	// 或者使用更简单的写法
	fmt.Print(string(buf[:n]))
}
