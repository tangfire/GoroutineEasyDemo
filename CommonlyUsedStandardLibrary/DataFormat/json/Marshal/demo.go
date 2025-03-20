package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Hobby string
}

type User struct {
	//"-"是忽略的意思
	Name  string `json:"-"`
	Hobby string `json:"hobby" `
}

func main() {
	p := Person{"5lmh.com", "女"}
	// 编码json
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))

	// 格式化输出
	b, err = json.MarshalIndent(p, "", "     ")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))

	fmt.Println("------------------")

	u := User{"tangfire", "男"}
	marshal, err := json.Marshal(u)
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(marshal))

	indent, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(indent))

}
