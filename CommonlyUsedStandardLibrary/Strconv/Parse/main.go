package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-2", 10, 64)
	u, _ := strconv.ParseUint("2", 10, 64)

	fmt.Println(b)
	fmt.Println(f)
	fmt.Println(i)
	fmt.Println(u)
}
