package main

import "fmt"

// 切片

func main() {
	var a1 []int
	var a2 []string
	fmt.Println(a1, a2)
	fmt.Println(a1 == nil) // 切片为空， nil表示未分配内存
	// 初始化
	a1 = []int{1, 2, 3}
	a2 = []string{"a", "b", "c"}
	fmt.Println(a1, a2)
	fmt.Println(a1 == nil)
	// 长度和容量
	fmt.Println("容量 ", cap(a1)) //3
	fmt.Println("长度 ", len(a1)) //3
	// 从一个数组中获得切片
	b1 := [...]int{1, 3, 4, 5, 6, 7}
	b2 := b1[0:4] // {1,3,4,5}
	b3 := b1[3:]  // {5,6,7}
	b4 := b1[:4]  // 同b2
	b5 := b1[:]   // 同b1

	fmt.Println(b2, b3, b4, b5)
	// 切片容量指向的底层数组第一个元素到最后一个元素的数量，即只能向右扩容
	fmt.Printf("len(b2):%d,cap(b2):%d,len(b3):%d,cap(b3):%d\n", len(b2), cap(b2), len(b3), cap(b3))
	// 切片再切片
	b6 := b3[1:]
	fmt.Printf("len(b6):%d,cap(b6):%d\n", len(b6), cap(b6))
	// 修改数组的值对切片的影响，切片是引用类型，都指向了底层的一个数组
	b1[1] = 300
	b1[5] = 700
	fmt.Println(b1, b2, b6) // 无论切割多少次，切片相同位置也受到了影响

	// make函数创造切片,切片的本质是一个框，框住了一块连续的内存
	c1 := make([]int, 5, 10) // 长度是5，容量为10
	fmt.Printf("c1-v-%v,c1-len-%d,c1-cap-%d\n", c1, len(c1), cap(c1))
	c2 := make([]int, 0, 10) // 长度是0，容量是10，c2的值为空
	fmt.Printf("c2-v-%v,c2-len-%d,c2-cap-%d\n", c2, len(c2), cap(c2))
	c3 := make([]int, 5) // 长度是5，容量是5，c3的值为[0,0,0,0,0]
	fmt.Printf("c3-v-%v,c3-len-%d,c3-cap-%d\n", c3, len(c3), cap(c3))

	// 切片比较
	var d1 []int
	fmt.Println(d1 == nil)
	var d2 = []int{}
	fmt.Println(d2 == nil) // 切片已经被初始化
	var d3 = make([]int, 0, 0)
	fmt.Println(d3 == nil) // 切片已经被初始化

	// 切片遍历
	for i := 0; i < len(a1); i++ {
		fmt.Println(i, a1[i])
	}
	fmt.Println("slice range loop")
	for k, v := range a1 {
		fmt.Println(k, v)
	}

}
