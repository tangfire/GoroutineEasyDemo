package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("Topgoer", "goer")) //true
	fmt.Println(strings.HasSuffix("Topgoer", "R"))    //false
	fmt.Println(strings.HasSuffix("Topgoer", "GOER")) //false
	fmt.Println(strings.HasSuffix("123456", "456"))   //true
	fmt.Println(strings.HasSuffix("Topgoer", ""))     //true
}
