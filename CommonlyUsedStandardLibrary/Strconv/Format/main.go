package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)

	fmt.Println(s1, s2, s3, s4)
}
