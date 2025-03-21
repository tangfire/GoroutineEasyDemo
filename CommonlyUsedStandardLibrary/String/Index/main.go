package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("Topgoer", "goer")) //true
	fmt.Println(strings.Index("Topgoer", "R"))    //false
	fmt.Println(strings.Index("Topgoer", "GOER")) //false
	fmt.Println(strings.Index("123-456", "-"))    //true
	fmt.Println(strings.Index("Topgoer", ""))     //true
}
