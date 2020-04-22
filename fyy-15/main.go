package main

// map

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

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

	// 遍历
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// 只遍历 key
	for k := range m2 {
		fmt.Println(k)
	}

	// 只遍历value
	for _, v := range m2 {
		fmt.Println(v)
	}

	// 删除key
	delete(m2, "海贼")

	fmt.Println(m2)

	// 删除不存在的 key
	delete(m2, "海贼王")

	//随机数种子，生成随机数map
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	fmt.Println(scoreMap)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	fmt.Println(scoreMap)

	// 无序遍历map，因为map是散列，需要借助切片实现
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}

	// 给map排序
	// 先创建空的切片
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 给切片排序
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	// map结合slice使用,切片和map在使用时，一定要进行初始化
	// 元素为map类型的切片
	var s1 = make([]map[string]int, 10, 10)
	mp := make(map[string]int, 1)
	mp["string"] = 10
	s1[0] = mp
	fmt.Println(s1)

	// 元素为切片类型的map
	var s2 = make(map[string][]int, 10)
	s2["slice0"] = []int{0, 1, 2}
	fmt.Println(s2)

	var s3 = make([]map[string]int, 1, 10)

	// map没有进行初始化
	//s3[0]["string"] = 199
	s3[0] = make(map[string]int, 1)
	s3[0]["string"] = 100
	fmt.Println(s3)

}
