package main

import "fmt"

// 构造函数
// 方法， 是作用于特定类型的函数
// 如果标识符的首字母的大写的，就表示对外部包可见的(暴露的，公有的)
// // 方法的定义
// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数){
// 	函数体
// }

type person struct {
	name string
	age  int
}

// 构造函数以 new 开头，返回值为结构体类型的名称，约定俗成
// 返回的是结构体还是结构体指针，取决于返回的结构体重量，
// 当需要大量拷贝结构体时，使用指针
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 方法命名规则，取类型的首个小写字母, 作为这个函数的接收者
// 值接收者，传递拷贝
func (p person) doSomething() {
	fmt.Printf("%s like watching TV\n", p.name)
	p.age++
}

// 指针接收者
// 适用类型
// 需要修改接收者中的值
// 接收着时拷贝代价比较大的对象
// 保证一致性，如果有某个方法使用了指针接收者，剩下的方法都要使用指针接收者
func (p *person) addAge() {
	p.age++
}

// 不能给别的包里面的类型添加方法， 只能自己给自己包里面的类型添加方法
// 给内置的int 类型添加一个方法
type myint int

func (m myint) phello() {
	fmt.Println("我是一个int类型的方法")
}

func main() {
	p1 := newPerson("a", 1)
	p2 := newPerson("b", 2)
	fmt.Println(p1, p2)
	p1.doSomething() // 使用person类型的方法
	p2.doSomething()
	p1.addAge()
	p2.addAge()
	fmt.Println(p1, p2)
	m := myint(10)
	m.phello()
	fmt.Println(m)
}
