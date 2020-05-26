package main

import "fmt"

// fmt占位符

// 数值类型
// func main() {
// 	n := 100
// 	fmt.Printf("%T\n", n) // 查看类型
// 	fmt.Printf("%v\n", n) // 打印值
// 	fmt.Printf("%b\n", n) // 二进制
// 	fmt.Printf("%d\n", n) // 十进制
// 	fmt.Printf("%o\n", n) // 八进制
// 	fmt.Printf("%x\n", n) // 十六进制
// 	s := string("abc")
// 	fmt.Printf("%s\n", s)  // 打印字符
// 	fmt.Printf("%#v\n", s) // 增加#，会在打印的时候将字符串用引号("")包围
// 	s1 := "ab次"
// 	c1 := 'a'
// 	c2 := 'b'
// 	c3 := '次'
// 	fmt.Println(s1, c1, c2, c3)
// }

// 通用占位符，精度
// func main() {
// 	var mp = make(map[string]int, 1)
// 	mp["zone"] = 1
// 	fmt.Printf("%v\n", mp)        //map[zone:1]
// 	fmt.Printf("%#v\n", mp)       //map[string]int{"zone":1}
// 	fmt.Printf("%d%%\n", 100)     // 100%
// 	// 宽度和精度，宽度表示打印出来的总长度
// 	fmt.Printf("%9.f\n", 12.345)  // (       12)
// 	fmt.Printf("%9.2f\n", 12.345) // (    12.35)
// 	fmt.Printf("%9f\n", 12.345)   // (12.345000)
// 	fmt.Printf("%.4f\n", 12.345)  // (12.3450)
// }

// // 用户输入,Scan
// func main() {
// 	var s string
// 	fmt.Scan(&s)
// 	fmt.Printf("用户输入的是：%s\n", s)
// }

// // 用户输入，Scanf
// func main() {
// 	var (
// 		name string
// 		age  int
// 	)
// 	fmt.Scanf("%s %d \n", &name, &age)
// 	fmt.Printf("用户输入的是：name: %s, age: %d\n", name, age)
// }

// 用户输入，Scanln
func main() {
	var (
		name string
		age  int
	)
	fmt.Scanln(&name, &age)
	fmt.Println("用户输入的是：", name, age)
}
