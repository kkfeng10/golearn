package main

// 整型

import "fmt"

func main() {
	// 类型推导，十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1) // 十进制转八进制
	fmt.Printf("%x\n", i1) // 十进制转十六进制
	fmt.Printf("%b\n", i1) // 十进制转二进制
	fmt.Println()

	// 八进制
	i2 := 077
	fmt.Printf("%o\n", i2)
	fmt.Printf("%d\n", i2) // 八进制转十进制
	fmt.Printf("%x\n", i2) // 八进制转十六进制
	fmt.Printf("%b\n", i2) // 八进制转二进制
	fmt.Println()

	// 十六进制
	i3 := 0x1234566
	fmt.Printf("%x\n", i3)
	fmt.Printf("%o\n", i3) /// 十六进制转八进制
	fmt.Printf("%d\n", i3) /// 十六进制转十进制
	fmt.Printf("%b\n", i3) /// 十六进制转二进制
	fmt.Println()

	// 二进制
	i6 := 0b1101
	fmt.Printf("%b\n", i6)
	fmt.Printf("%o\n", i6) /// 二进制转八进制
	fmt.Printf("%d\n", i6) /// 二进制转十进制
	fmt.Printf("%x\n", i6) /// 二进制转十六进制
	fmt.Println()

	// 查看变量的类型
	fmt.Printf("%T\n", i3)
	i7 := "abc"
	fmt.Printf("%T\n", i7)
	var i8 int8 = 9
	fmt.Printf("%T|%d\n", i8, i8)

}
