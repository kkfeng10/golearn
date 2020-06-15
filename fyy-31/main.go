package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// 匿名结构体字段，结构体嵌套，json序列化与反序列化

type address struct {
	province string
	city     string
}

// 使用 address 匿名结构体，可以在结构体变量中直接调用，而不需要指定结构体名称，当存在同名字段时，为防止冲突，是需要指定的
// 如果需要使用json进行序列化，字段必须使用首字母大写， 使用反引号进行json反序列化时的改写
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	address
}

// 匿名结构体的方法也是可以在结构体变量中直接调用的
func (a address) registry() {
	fmt.Println("registry to a new place")
}

func main() {
	p1 := &person{
		Name: "x",
		Age:  17,
		address: address{
			province: "js",
			city:     "nj",
		},
	}
	fmt.Println(p1)
	p1.registry()
	j1, _ := json.Marshal(p1)
	fmt.Println(string(j1))

	p2 := &person{}
	json.Unmarshal([]byte(j1), p2)
	fmt.Printf("%#v\n", p2)
	fmt.Println(p2)
	fmt.Println(math.Pi)
}
