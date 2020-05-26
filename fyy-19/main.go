package main

import "fmt"

// defer面试题
// 秘诀： defer 虽然在最后执行，但是在注册的时候参数状态是不会在后面发生改变的
// 调用defer的过程其实就是入栈出栈的过程

func deferPractis(x string, a int, b int) (ret int) {
	ret = a + b
	fmt.Println(x, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer deferPractis("1", a, deferPractis("10", a, b))
	a = 0
	defer deferPractis("2", a, deferPractis("20", a, b))
	b = 1
}

// run: 10 1 2 3
// deferPractis("1",1,3)
// run: 20 0 2 2
// deferPractis("2",0,2)
// deferPractis("2",0,2) >> run: 2 0 2 2
// deferPractis("1",1,3) >> run: 1 1 3 4

// 打印输出结果
// 10 1 2 3
// 20 0 2 2
// 2 0 2 2
// 1 1 3 4
