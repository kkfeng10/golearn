package main

import "fmt"

// 接口
// 因为接口可以存任意类型，所以是动态类型，值也是动态值，类型和值是分开存储的

// 定义接口类型，使用该接口则必须要有speak方法
type speaker interface {
	speak()
}

//  定义一个cat 类型的结构体，并在下方实现了 speak 方法
type cat struct {
	name string
}
type dog struct {
	name string
}
type person struct{}

func (c cat) speak() {
	fmt.Printf("my name is %s,miu~~~\n", c.name)
}

func (d dog) speak() {
	fmt.Printf("my name is %s,wang~~~\n", d.name)
}

func (p person) speak() {
	fmt.Println("yiya~~~")
}

func da(x speaker) {
	x.speak()
}

func main() {
	// 定义cat类型的变量c
	var c = cat{
		name: "abc",
	}
	// 定义dog类型的变量d，进行了初始化操作
	d := dog{
		name: "bcd",
	}
	// 定义person类型的变量p, 未进行初始化
	var p person

	da(c)
	da(d)
	da(p)

}
