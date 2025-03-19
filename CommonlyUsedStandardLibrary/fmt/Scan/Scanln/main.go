package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("name:%s age:%d married:%t \n", name, age, married)
}
