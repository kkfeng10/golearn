package main

import "fmt"

// 函数存在的意义?
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中，给他起个名字，每次用到它的时候用函数名就可以
// 返回值一定要声明其类型
// Go语言中函数没有默认参数这个概念

// 定义函数,有返回值,提前命名，声明了一个变量ret
func sum1(a int, b int) (ret int) {
	ret = a + b + 2
	return a + b
}

// 没有返回值
func sum2(a int, b int) {
	fmt.Println(a + b)
}

// 定义函数，有返回值，提前命名，返回时可省略变量
func sum3(a int, b int) (ret int) {
	ret = a + b
	return
}

// 多个返回值
func sum4(a int, b int) (int, int) {
	return a + b, a - b
}

// 参数类型简写, 当参数中连续多个参数的类型一致时，我们可以将非最后一个参数的类型省略
func sum5(a, b int) int {
	return a + b
}

func sum6(a, b int, c, d, string, e, f bool) string {
	return "abc"
}

// 可变参数, b是切片类型
func sum7(a int, b ...int) string {
	fmt.Printf("sum7,b type is  %T\n", b) // []int 类型
	return "abc"
}

func main() {
	r1 := sum1(1, 2)
	fmt.Println(r1)
	r2 := sum3(3, 4)
	fmt.Println(r2)
	fmt.Println(sum5(7, 8))
	fmt.Println(sum7(10, 11))

	a1 := []int{1, 2, 3}
	a2 := a1
	// var a3 []int   // 未初始化，a3 为 nil
	// a3 := make([]int, 0, 3) // 虽然初始化了容量，但是没有允许拷贝的位置空间
	a3 := make([]int, 3, 3)
	copy(a3, a1)
	a1[1] = 200             // 修改 底层数组的值
	fmt.Println(a1, a2, a3) // [1 2 3] [1 2 3] []

}
