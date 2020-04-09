package main

import "fmt"

// 数组

// 存放元素的容器
// 必须指定存放元素的类型和长度

func main() {
	// 数组的长度是数组类型的一部分,所以不同的长度之间无法比较
	var a1 [3]bool // [true false false]
	var a2 [5]string
	var a3 [4]int
	fmt.Printf("a1 type: %T | a2 type: %T \n", a1, a2)

	// 数组不初始化，默认元素都是零值(布尔值:false, 整型和浮点型都是0， 字符串: "")
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	// 数组初始化
	// 1. 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2. 初始化方式2,根据初始值自动进行数组长度的推断，这也是使用方法最多的一种
	a4 := [...]int{1, 2, 3, 4, 5, 5}
	fmt.Println(a4)
	// 3. 初始化方式3，根据索引来进行初始化
	a5 := [5]int{0: 1, 4: 2}
	fmt.Println(a5)

	// 遍历
	cities := [...]string{"纽约", "芝加哥", "波士顿"}
	// 1. 根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println("intex:", i, cities[i])
	}
	// 2. for range 遍历
	for i, v := range cities {
		fmt.Printf("range:%d %s\n", i, v)
	}

	// 多维数据
	// 二维数组的声明
	var a6 [3][2]int
	a6 = [3][2]int{
		[2]int{1, 2},
		[2]int{2, 3},
		[2]int{3, 4},
	}
	fmt.Println(a6)
	fmt.Printf("a6 type is %T \n", a6)
	// 二维数组的遍历
	for _, i := range a6 {
		fmt.Println(i)
		fmt.Printf("a6 sub arrray type is %T \n", i)
		for _, j := range i {
			fmt.Println(j)
		}
	}

	// 数组是值类型
	a7 := [...]int{1, 3, 5}
	a8 := a7
	a8[0] = 2
	fmt.Println(a7, a8)
}
