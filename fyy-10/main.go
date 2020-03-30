package main

import "fmt"

// switch
// 简化大量的判断 if...(else if...)*n else...
// goto

func main() {
	// 作用域在函数内
	n := 10
	switch n {
	case 1:
		fmt.Println("one")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("unknow number", n) //这里会主动在不同元素之间添加空格打印
	}

	// 作用域在switch内
	switch m := 10; m {
	case 1, 3, 5, 7, 9: // case可用于选择列表
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println("unknow number", m) //这里会主动在不同元素之间添加空格打印
	}

	// fmt.Println(m) //作用域只在switch内

	// 分支使用表达式,switch后不能直接跟变量,但是可以先给变量赋值
	switch p := 10; {
	case p > 7 && p < 12:
		fmt.Println("in 7 to 12")
	case p >= 12:
		fmt.Println("above 12")
	default:
		fmt.Println("unknow number", p)
	}

	// goto，不推荐使用
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if i == 3 {
				goto GOTOLABEL
			}
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

GOTOLABEL:
	fmt.Println("over")

	// goto的另一种替代方案,使用 flag 进行替代
	println("over?")

	flag := false
	for i := 1; i < 10; i++ {
		if flag == true {
			break
		}
		for j := 1; j <= i; j++ {
			if i == 3 {
				flag = true
				break
			}
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

}
