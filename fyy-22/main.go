package main

import "fmt"

// 结构体
// 自定义类型和类型别名, 使用type 关键字来进行构造

type myInt int     // 基于 int 类型造一个 myInt 类型,
type yourInt = int // 类型别名， 比如 rune 代表的是 int32  , 使用%T打印出来的是原类型

func main() {
	var p myInt
	p = 100
	fmt.Printf("type: %T, value: %v\n", p, p)

	var q yourInt
	q = 200
	fmt.Printf("type: %T, value: %v\n", q, q)

	var r rune
	r = '习' // 使用 int32 来表示字符
	fmt.Printf("type: %T, value: %v\n", r, r)

	for _, v := range "好好学习" {
		fmt.Println(string(v)) // 强制转换成 string 类型
		fmt.Printf("%T\n", v)  // 这里打印出来的是rune类型(int32)
	}

	var s string
	s = "好好学习"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
		fmt.Printf("%T\n", s[i])
	}

}
