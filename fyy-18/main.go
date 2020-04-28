package main

import "fmt"

// 匿名函数 , 一般用在函数内部

// 闭包， 当一个函数的参数类型不匹配的时候，可以先封装一层能够匹配的参数类型，然后调用时执行包内的函数
// 特点|目的： 能包含外部作用域的变量

// 闭包1,需要传入func()类型的参数
func b1(f func()) {
	f()
}

// 闭包2，需要传入 int int 类型的参数，没有返回值，但是想要使用f1函数,并且在其中执行
func f2(x, y int) {
	fmt.Println("in f2")
	fmt.Println(x + y)
}

// 闭包3，借助此闭包来实现 f2 的需求
func f3(f func(int, int), m, n int) func() {
	tmp := func() {
		f(m, n)
	}
	return tmp
}

func main() {
	// 匿名函数赋值给变量
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(1, 2)

	// 直接调用匿名函数
	func(x, y int) {
		fmt.Println(x * y)
	}(2, 3)

	b1(f3(f2, 100, 300)) // 层层拨开
}
