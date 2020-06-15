package calc

import "fmt"

// init函数，导入此包时自动执行
func init() {
	fmt.Println("calc包中的 init 函数")
}

// Add 对外暴露的函数，首字母要大写, 且必须以函数名开头写上注释
func Add(x, y int) int {
	return x + y
}
