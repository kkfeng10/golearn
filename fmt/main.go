package main

import "fmt"

// fmt占位符

func main() {
	n := 100
	fmt.Printf("%T\n", n) // 查看类型
	fmt.Printf("%v\n", n) // 打印值
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%d\n", n) // 十进制
	fmt.Printf("%o\n", n) // 八进制
	fmt.Printf("%x\n", n) // 十六进制
	s := string("abc")
	fmt.Printf("%s\n", s)  // 打印字符
	fmt.Printf("%#v\n", s) // 增加#，会在打印的时候将字符串用引号("")包围
	s1 := "ab次"
	c1 := 'a'
	c2 := 'b'
	c3 := '次'
	fmt.Println(s1, c1, c2, c3)
}
