package main

import "fmt"

// 函数作为参数和返回值

func f1() {
	fmt.Println("f1")
}

func f2() int {
	fmt.Println("f2")
	return 1
}

func f3(x, y int) int {
	return x + y
}

// 函数作为参数
func f4(x func() int) {
	fmt.Println("函数做参数")
}

// 函数作参数，且函数作为返回值
func f5(x func(int, int) int) func(int, int) int {
	return f3
}

func main() {
	a1 := f1 // 无返回值的 func() 类型
	a2 := f2 // 有返回值的 func() int 类型
	a3 := f3 // 带参数的有返回值的 func(int,int) int 类型
	fmt.Printf("a1 type is %T, a2 type is %T, a3 type is %T", a1, a2, a3)
}
