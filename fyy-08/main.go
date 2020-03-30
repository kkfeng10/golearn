package main

import "fmt"

// byte和rune类型
// 类型转换

func main() {
	// byte型遍历字符串
	s := "hello世界"
	fmt.Printf("len: %v\n", len(s)) //一个中文字符的长度是3
	for i := 0; i < len(s); i++ {   //byte类型处理方式
		fmt.Printf("%T|%v|%c\n", s[i], s[i], s[i]) // byte类型是uint8，非英文字符会打印出乱码
	}
	fmt.Println()
	// rune类型遍历字符串
	for _, r := range s {
		fmt.Printf("%T|%v|%c\n", r, r, r) // rune类型是int32
	}
	fmt.Println()
	// 单引号和双引号的类型区别
	s1 := "好"
	c1 := '好'                                               //  int32
	c2 := 'h'                                               // int32, 注意, 单引号声明的都是int32类型
	c3 := byte('h')                                         // int32 强制转换成 uint8 类型
	fmt.Printf("s1:%T c1:%T c2:%T c3:%T\n", s1, c1, c2, c3) // s1是string类型，c1是 int32类型，也就是rune类型
	fmt.Println()
	// 字符串转换
	s2 := "好吗"
	// s2[1] = "的" //不允许这种转换，字符串类型并不是数组类型
	s3 := []rune(s2) // 将字符串强制转换成rune类型的切片
	// s3[1] = "的" //不允许使用字符串给rune类型的字符赋值
	s3[1] = '的'
	fmt.Println(string(s3)) // 打印的的时候需要将rune类型的切片转换成字符串
	fmt.Println()

	// 整型和浮点型转换
	n := 10
	f := float64(n)
	fmt.Printf("n:%T f:%T\n", n, f)

	// 统计汉子的数量
	fun := "hello世界,你好世界"
	count := 0
	for _, w := range fun {
		if len(string(w)) == 3 {
			count++
		}
	}
	fmt.Println(count)
}
