package main

import "fmt"

// 函数存在的意义?
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中，给他起个名字，每次用到它的时候用函数名就可以
// 返回值一定要声明其类型
// Go语言中函数没有默认参数这个概念

// defer 函数内延迟执行

// 注意： 函数中 return语句在底层并不是原子操作，它分为赋值和ret指令2个步骤
// 而 defer 语句执行的时机就在返回值赋值操作后，ret指令执行前

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

// 可变参数, b是切片类型，表示若干个int类型，只能最后一个写
func sum7(a int, b ...int) string {
	fmt.Printf("sum7,b type is  %T\n", b) // []int 类型
	return "abc"
}

// defer 延迟执行函数，在函数将要推出的时候执行， 多用于函数结束之前释放资源
// 一个函数中可以有多个defer语句，按照后进先出的顺序延迟执行
func deferdemon() {
	fmt.Println("打开文件")
	defer fmt.Println("开始关闭文件") // 最先放进去的，最后执行
	defer fmt.Println("准备关闭文件") // 最后放进去的，最先执行
	fmt.Println("开始执行")
}

// 使用未命名的返回值，返回值为1，修改的是x ，不是返回值
func defer1() int {
	x := 1
	defer func() {
		x++
	}()
	return x
}

// 使用有命名的返回值，返回值为6，return 先赋值给x, defer延迟执行函数修改x的值，最后执行RET 返回x
func defer2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// 使用有命名的返回值，返回值6, return 先将x赋值给y, defer 延迟修改x的值，但是不影响y的返回值
func defer3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 使用有命名的返回值，返回值为5，return 先将x赋值为5， 然后将x作为参数传递给defer函数，defer函数修改的是作用域在defer函数体内形参的值，不会对原本x造成影响
func defer4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// 返回5，defer函数没有返回值接收
func defer5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5
}

// 返回6，传递一个指针，defer函数直接到内存中修改返回值
func defer6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5
}

func main() {
	r1 := sum1(1, 2)
	fmt.Println(r1)
	r2 := sum3(3, 4)
	fmt.Println(r2)
	fmt.Println(sum5(7, 8))
	fmt.Println(sum7(10, 11))

	// 无法再命名函数中声明命名函数
	// func sum8(x,y int) int{
	// 	return x+y
	// }

	deferdemon()
	fmt.Println(defer1())
	fmt.Println(defer2())
	fmt.Println(defer5())
	fmt.Println(defer6())

}
