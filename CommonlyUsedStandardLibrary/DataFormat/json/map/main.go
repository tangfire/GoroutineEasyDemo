package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	student := make(map[string]interface{})
	student["name"] = "5lmh.com"
	student["age"] = 18
	student["sex"] = "man"
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	fmt.Println("-----------------")
	fmt.Println(string(b))
}
