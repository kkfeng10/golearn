package main

import (
	"fmt"
	"strings"
)

// 匿名函数 , 一般用在函数内部

// 闭包， 当一个函数的参数类型不匹配的时候，可以先封装一层能够匹配的参数类型，然后调用时执行包内的函数
// 特点|目的： 能包含外部作用域的变量
// 闭包利用了函数是可以作为返回值这个特点，函数内部查找变量的顺序，先在自己内部找，找不到往外层找

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

// 闭包4，使用外部作用域的一个变量
func f4(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包5， 实际应用例子，字符串后缀处理
func f5(subfix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, subfix) {
			return name + subfix
		}
		return name
	}
}

// 闭包6， 实际应用例子， 数值计算
func f6(base int) (func(int) int, func(int) int) {
	add := func(x int) int {
		base += x
		return base
	}

	cut := func(x int) int {
		base -= x
		return base
	}

	return add, cut

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

	ret := f4(100)
	ret1 := ret(200)
	fmt.Println("ret1: ", ret1)

	jpgSub := f5(".jpg") // 先将 .jpg 传入到闭包中，缓存变量
	pngSub := f5(".png") // 先将 .png 传入到闭包中，缓存变量

	fmt.Println(jpgSub("abc.txt")) // 使用缓存变量进行运算
	fmt.Println(pngSub("aaa.png"))

	add, cut := f6(10) // 缓存变量 10
	// 由于共享缓存变量，所以，下面的计算过程中 base 变量的值会一直存在
	fmt.Println(add(1), cut(4)) // 11,7
	fmt.Println(add(2), cut(5)) // 9,4
	fmt.Println(add(3), cut(6)) // 7,1
}
