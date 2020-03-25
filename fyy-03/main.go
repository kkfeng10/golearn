package main

import "fmt"

// 常量

const pi = 3.1415926

// 批量声明常量
const (
	statusOk = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
	// n1=n2=n3=100
)

// iota ,const关键字出现时被重置为0，const中每新增一行常量声明，将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
// 使用iota能简化定义，在定义枚举时很有用
const (
	a1 = iota //0
	a2        // = iota =1
	a3        // = iota =2
)

// 加入匿名变量
const (
	b1 = iota //0
	b2        //1
	_
	b3 //3
)

// 插队
const (
	c1 = iota //0
	c2 = 100
	c3        //100
	c4 = iota //3
	c5        //4
)

// 多个iota声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1:1 ,d2:2
	d3, d4 = iota + 1, iota + 2 // d3:2 ,d3:3
)

// 定义数量级,iota的一个应用
const (
	// << 左移，比如 1 << 2 ,代表 10 （二进制）2的1次方
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(pi)
	fmt.Println(n1, n2, n3)
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)
	fmt.Println("c5:", c5)
	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)
}
