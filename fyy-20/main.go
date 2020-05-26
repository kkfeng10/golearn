package main

// panic,recover
// 异常处理
// recover 会尝试修复panic的错误，但是最好不要滥用

import "fmt"

// panic结合 defer使用
// func funcA() {
// 	// 数据库连接应用场景
// 	fmt.Println("创建数据库连接")
// 	defer func() {
// 		fmt.Println("释放数据库连接")
// 	}()
// 	panic("数据库操作崩溃")
// }


// panic,recover 结合 defer 使用, defer 一定要在Panic 之前使用
func funcA() {
	// 数据库连接应用场景
	fmt.Println("创建数据库连接")
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("数据库操作崩溃")
}

func main() {
	funcA()
	fmt.Println("程序退出")
}
