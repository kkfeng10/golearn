package main

import "fmt"

// 推荐使用驼峰命名
// startEnd

// 批量声明
var (
	s1 string // ""
	s2 int    // 0
	s3 bool   // false
)

func foo() (int, string) {
	return 10, "yff"
}

func main() {
	s1 = "name"
	s2 = 1
	s3 = true
	// Go中变量声明了必须使用，不使用无法编译(后面的版本应该已经改动了)
	// Print, 只打印不换行
	// Println, 打印并换行、
	// Printf, 格式化打印，需要自己换行
	fmt.Print(s3)
	fmt.Println()
	fmt.Println(s2)
	fmt.Printf("name:%s\n", s1)
	fmt.Printf("name:%s\n", s1)

	// 声明变量的同时赋值
	var s4 string = "yyf"
	var s7 = string("yyf")
	// 类型推导(根据值来判断)
	var s5 = "yyf5"
	// 简短变量声明,只能在函数里面用
	s6 := "yyf6"
	s8 := string("yyf8")
	fmt.Println(s4, s5, s6, s7, s8)
	// 较为常用的，声明全局变量，用批量声明
	// 声明局部变量，用 := 声明
	// 有较为特殊的场景，用 var s5 = "yyf5" 声明

	// 匿名变量，不占用命名空间，不会分配内存，不存在重复声明，占位，表示忽略某个值
	x, _ := foo()
	// 同一个作用域中不能重复声明同名变量
	// s1 := 123
	fmt.Println(x)
}
