package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 1, 判断字符串中汉字的数量
// 2, 判断字符串中单词的出现个数
// 3, 回文判断，字符串从左往右读和从右往左读时一样的

func main() {
	// 1. 判断字符串中汉字的数量
	s := "abc你好"
	var count int
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)

	// 2, 判断字符串中单词的出现个数
	p := "How do you do"

	// 字符串切割
	sp := strings.Split(p, " ")

	cs := make(map[string]int, 10)

	for _, char := range sp {
		_, ok := cs[char]
		if !ok {
			cs[char] = 1
		} else {
			cs[char]++
		}
	}

	for key, value := range cs {
		fmt.Println(key, value)
	}

	//   3, 回文判断，字符串从左往右读和从右往左读时一样的
	ss := "黄山落叶松叶落山黄"
	// 直接使用for循环打印，其结果是uint8类型的数字, 中文字符占3个字节，需要用32为utf8来表示
	// for i := 0; i < len(ss); i++ {
	// 	fmt.Println(ss[i])
	// 	fmt.Printf("%T\n", ss[i])
	// }

	// 使用切片来处理
	rs := make([]rune, 0, len(ss))
	for _, c := range ss {
		rs = append(rs, c)
	}

	isBlang := true

	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-1-i] {
			isBlang = false
			break
		}
	}
	if isBlang {
		fmt.Println("是回文")
	} else {
		fmt.Println("不是回文")
	}

}
