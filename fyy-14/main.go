package main

// 指针，new, make

import "fmt"

func main() {
	n := 18
	// 1. 取地址 &
	fmt.Println(&n)
	fmt.Printf("%T \n", &n)
	// 2. 根据地址取值 *
	m := &n
	fmt.Println(*m)

	// 声明一个int类型的指针，注意，此时为空指针
	var a1 *int
	//fmt.Println(*a1) // 空指针无法寻址，所以没有找到指定的值
	fmt.Println(a1) // 打印的值为nil

	// 使用new创建一块内存,一般用来给基本数据类型申请内存，string/int ..., 返回的是对应类型的指针， *string, *int
	a2 := new(int)
	// 找到内存地址赋值
	*a2 = 100
	fmt.Println(a2, *a2)

	// make, 内存分配，区别new, 它只用于slice, map 以及 chan 的内存创建，返回的对应三个类型本身

}
