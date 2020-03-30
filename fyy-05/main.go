package main

import (
	"fmt"
	"math"
)

// 浮点数

func main() {
	f1 := math.MaxFloat32
	fmt.Printf("%f\n", f1)
	f2 := 123.123123
	fmt.Printf("%T\n", f2) //默认小数都是float64类型
	f1 = f2
	fmt.Printf("%f\n", f1) //同为float64类型可以赋值操作
	f3 := float32(456.456456)
	fmt.Printf("%T\n", f3)
	// f1 = f3  //不同浮点数类型无法进行赋值操作
}
