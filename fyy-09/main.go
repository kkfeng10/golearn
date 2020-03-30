package main

import "fmt"

// if流程控制
// for循环

func main() {
	// 变量在整个函数的作用域判断
	day := 5
	if day <= 4 {
		fmt.Println("busy")
	} else if day >= 6 {
		fmt.Println("relax")
	} else {
		fmt.Println("easy")
	}
	// 变量只在if当中生效
	if age := 16; age >= 18 {
		fmt.Println("祝您游戏愉快")
	} else {
		fmt.Println("未满十八岁禁止入内")
	}
	// fmt.Println(age) //变量作用域仅在if当中生效

	// for循环基本格式， 最常用的
	for i := 0; i < 3; i++ {
		fmt.Print(i)
	}
	//fmt.Println(i) // 变量i作用域仅在for循环体中生效
	fmt.Println()

	// for循环变种一,使用已经定义好的变量,省略初始变量
	i := 5
	for ; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	fmt.Println(i) // 此时i的值已经是10,因为每次循环体结束后都要+1

	// for循环变种二，省略初始条件和结束时的条件
	for i < 15 {
		fmt.Print(i)
		fmt.Print(" ")
		i++
	}
	fmt.Println()

	// 死循环
	// for {
	// 	fmt.Print('.')
	// }

	// for range （键值循环），使用较多
	for x, y := range "hello世界" {
		fmt.Printf("%d,%c\n", x, y)
	}

	// 99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
	// 99倒乘法表
	for i := 8; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
	fmt.Println()

	// break跳出标签，知道是干什么的就行了，基本上用不到
BREAKTAG:
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if i == 3 {
				break BREAKTAG
			}
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

	fmt.Println()
	// continue跳转标签，知道是干什么的就行了，基本上用不到
CONTINUETAG:
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 3 {
				fmt.Println()
				continue CONTINUETAG
			}
			fmt.Printf("%dx%d=%d\t", i, j, i*j)
		}
	}
	fmt.Println()
}
