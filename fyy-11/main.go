package main

import "fmt"

// 运算符

func main() {
	// 算数运算符
	var (
		a = 5
		b = 3
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++ // 单独的语句，不能放在=的右边赋值，其他语言中，相当于一个值，可以赋值给其他变量
	b--

	// 关系运算符 逻辑运算符
	age := 28
	if age >= 18 && age <= 60 { // 逻辑与
		fmt.Println("年轻人")
	} else {
		fmt.Println("老年人")
	}

	if age >= 18 || age <= 60 { // 逻辑或
		fmt.Println("不用上班真好")
	} else {
		fmt.Println("赶紧上班去")
	}

	isWorking := false
	fmt.Println(!isWorking) // 逻辑非

	// 位运算: 针对二进制

	// & 按位与，两位均为1才位1
	fmt.Println(5 & 2) // 101 & 010 = 000
	// | 按位或，两位有一位为1则为1
	fmt.Println(5 | 2) // 101 | 010 = 111
	// ^ 异或， 两位不一样则为1
	fmt.Println(5 ^ 2) // 101 ^ 010 = 111
	// << 左移指定位数
	fmt.Println(5 << 2) // 101 << 2 = 10100
	// >> 右移指定位数
	fmt.Println(5 >> 2) // 101 >> 2 = 1
	// 左移时需要注意变量的范围，超出范围，可以运行但只会按照变量所支持的最大位数取值
	var m = int8(1)
	fmt.Println(m << 10) // 超出范围，只取变量的后8位

	// 赋值运算符，用来给变量赋值
	p := 10
	p += 2 // p+=1 使用 p++ 替代
	fmt.Println(p)
	p *= 2
	p /= 2
	p -= 2
	p %= 2

}
