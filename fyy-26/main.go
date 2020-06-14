package main

import "fmt"

// 接口的定义和实现
// 接口的类型和值是分开存储的，因为接口可以存任意类型，所以是存储的是动态类型，值也是动态值
// 必须实现了接口所有的方法才能够使用该接口类型
// 注意：如果使用指针接收者实现接口的方法，那么给接口类型变量赋值时也必须是指针类型的变量

// type animal interface {
// 	move()
// 	eat(string)
// }

// 使用接口嵌套来定义animal接口
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type chicken struct {
	name string
	feet int
}

type cat struct {
	name string
	feet int
}

// func (ch chicken) move() {
// 	fmt.Println("咯咯哒")
// }

// func (ch chicken) eat(food string) {
// 	fmt.Printf("小鸡啄%s\n", food)
// }

func (ch *chicken) move() {
	fmt.Println("咯咯哒")
}

func (ch *chicken) eat(food string) {
	fmt.Printf("小鸡啄%s\n", food)
}

func (c cat) move() {
	fmt.Println("喵")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {
	kfc := &chicken{
		name: "战斗鸡",
		feet: 2,
	}
	// kfc := chicken{
	// 	name: "战斗鸡",
	// 	feet: 2,
	// }
	miu := cat{
		name: "千岁岁",
		feet: 4,
	}
	// 定义接口类型变量
	var ani animal
	// 赋值chicken 类型给接口变量
	ani = kfc
	ani.eat("米")
	fmt.Println(ani)        //{战斗鸡 2} 打印的是接口的值，接口的值是动态的
	fmt.Printf("%T\n", ani) //main.chicken 打印的是接口的类型，类型也是动态的
	ani = miu
	miu.eat("小鱼干")
	fmt.Println(ani)        //{千岁岁 4}
	fmt.Printf("%T\n", ani) //main.cat
}
