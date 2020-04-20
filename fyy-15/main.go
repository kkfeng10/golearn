package main

// map

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1) // 还未初始化，map[]

	m2 := make(map[string]int, 10) // 需要预估好map容量，避免在程序运行期间再动态扩容
	m2["海贼"] = 15
	m2["火影"] = 90
	fmt.Println(m2) //map[海贼:15 火影:90]

	// 当取值无法判断key是否存在时时，使用ok进行接收判断（接收返回的一个bool值），约定好的
	value, ok := m2["柯南"]
	if !ok {
		fmt.Println("没有这个key")
	} else {
		fmt.Println(value)
	}

	// 如果不存在这个key, 则拿到对应类型的零值
	fmt.Println(value)

}
