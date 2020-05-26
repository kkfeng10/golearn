package main

import "fmt"

// 递归
// 适合处理的问题相同，并且问题的规模越来越小的场景
// 一定有一个明确的推出条件

// 计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 有n个台阶，一次上1步，一次上2步，一共有多少种走法
func f1(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f1(n-1) + f1(n-2)
}

func main() {
	p := f(5)
	fmt.Println(p)
	p1 := f1(5)
	fmt.Println(p1)
}
