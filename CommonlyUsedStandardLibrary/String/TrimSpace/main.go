package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimSpace(" I love my country "))
	fmt.Println(strings.TrimSpace(" \t\n  I love my country \t\n "))
	fmt.Println(strings.TrimSpace(" \t\n\r\x0BI love my country\t\n "))
}
