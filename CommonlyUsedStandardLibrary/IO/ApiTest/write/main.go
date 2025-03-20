package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件（使用更安全的权限0644）
	file, err := os.Create("./text.txt")
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer func() {
		// 确保关闭前同步数据到磁盘
		if err := file.Sync(); err != nil {
			fmt.Println("同步数据失败:", err)
		}
		if err := file.Close(); err != nil {
			fmt.Println("关闭文件失败:", err)
		}
	}()

	for i := 0; i < 5; i++ {
		// 写入字符串并检查错误
		if _, err := file.WriteString("hello world\n"); err != nil {
			fmt.Println("写入字符串失败:", err)
			return
		}

		// 写入字节数据并检查错误
		if _, err := file.Write([]byte("cd\n")); err != nil {
			fmt.Println("写入字节失败:", err)
			return
		}
	}

	fmt.Println("文件写入完成，请检查 text.txt")
}
