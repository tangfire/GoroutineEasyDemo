package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(time.Hour)
	fmt.Println(later)
	before := now.Add(-time.Hour)
	fmt.Println(before)

}
