package main

import (
	"fmt"
	"sort"
)

// 切片进阶

func main() {
	s1 := []string{"a", "b", "c", "d"}
	fmt.Printf("s1-v-%v,s1-len-%d,s2-cap-%d\n", s1, len(s1), cap(s1))

	// append添加元素,必须使用原来的变量接收返回值
	// 扩容规则，和使用的数据类型有关系
	// 如果新申请容量大于2倍旧容量，则最终容量就是新申请的容量
	// 否则，如果旧切片的长度小于1024, 则最终容量就是旧容量的两倍，即newcap=doublecap
	// 否则，如果旧切片长度大于等于1024, 则最终容量 从 旧容量 开始循环增加原来的 1/4 , 直到最终容量大于等于新申请的容量，即 newcap >= cap
	s1 = append(s1, "e") // 原来的底层数组放不下的时候，go会将底层数组更换一个位置
	fmt.Printf("s1-v-%v,s1-len-%d,s2-cap-%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "f", "g") // 原来的底层数组放不下的时候，go会将底层数组更换一个位置
	fmt.Printf("s1-v-%v,s1-len-%d,s2-cap-%d\n", s1, len(s1), cap(s1))
	ss := []string{"h", "i", "j"}
	s1 = append(s1, ss...) // ... 表示拆开
	fmt.Printf("s1-v-%v,s1-len-%d,s2-cap-%d\n", s1, len(s1), cap(s1))

	// 复制 copy
	a1 := []int{1, 2, 3}
	a2 := a1
	a3 := make([]int, 3, 3)
	// 将 a1 的底层数组元素拷贝到 a3 中，相当于把切片所指向的数组开辟了一块内存出来，不与原来的放在一起
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	// 修改底层数组，使用切片直接等号赋值，所指向的将是同一个底层数组
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 删除元素，切片一定是一段连续的内存,将会修改底层数组，存在覆盖元素的情况
	// 切片不保存值
	c1 := [...]int{1, 2, 3, 4}
	c2 := c1[:]
	fmt.Println(c1, c2)
	// 使用append拼接方法，删除第2个元素
	c2 = append(c2[:1], c2[2:]...)
	// c1 会在末尾追加一个元素，与原来的末尾元素相同，长度不变
	fmt.Println(c1, c2) // c1=[1 3 4 4] c2=[1 3 4]

	// sort.Ints方法排序元素是int类型的数组
	c3 := [...]int{3, 1, 2}
	sort.Ints(c3[:])
	fmt.Println("排序后: ", c3)
	// 练习题
	// 考验对切片初始化以及append操作的理解
	d1 := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		d1 = append(d1, i)
	}
	fmt.Println(d1) //[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]

}
