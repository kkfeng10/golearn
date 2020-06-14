package main

import (
	"fmt"
	calc "github/golearn/calc01" // 使用别名，否则包的名字就是目录的名字，二者应该是一致的，推荐目录名和包名是一致的
)

func init() {
	fmt.Println("main函数的init函数，导入包执行的init函数在我前面")
}

func main() {
	fmt.Println(calc.Add(10, 20))
}
