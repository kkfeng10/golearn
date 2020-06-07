package main

import "fmt"

// 结构体 匿名结构体

// 声明结构体方法
// type 类型名 struct {
// 	字段名 字段类型
// 	字段名 字段类型
// 	...
// }

// 声明结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

type pask struct {
	a int8
	b int8
	c int8
}

type oask struct {
	a int8
	b int8
	c string
	d string
}

func f1(x person) {
	x.name = "yyff"
}

// 传入指针类型的结构体
func f2(x *person) {
	(*x).name = "yyff"
}

func main() {
	// 声明一个person类型的变量, 结构体实例化
	var p person
	p.name = "ffyy"
	p.age = 16
	p.gender = "man"
	p.hobby = []string{"a", "b"} // 切片初始化
	fmt.Printf("%T\n", p)        //  main.person 类型
	fmt.Println(p)               // 打印

	// 匿名结构体, 多用于临时场景
	var s struct {
		s1 string
		s2 int
	}
	s.s1 = "s1"
	s.s2 = 1
	fmt.Printf("type: %T , values: %v\n", s, s)

	f1(p)
	fmt.Println(p.name)

	f2(&p)
	fmt.Println(p.name)

	// 声明结构体指针
	np := new(person)
	(*np).name = "gaey" // 也可以使用语法糖， np.name = "gaey" , 以后尽量使用语法糖的模式

	f2(np)
	fmt.Println(np.name)
	fmt.Printf("%T,%p\n", np, np)

	// 结构体初始化

	// k-v 初始化，并返回指针
	p2 := &person{
		name:   "hero",
		age:    16,
		gender: "man",
		hobby:  []string{"b", "c"},
	}

	p2.name = "hera"
	fmt.Printf("%v\n", *p2)

	// 列表初始化，并返回指针
	p3 := &person{
		"hero",
		16,
		"man",
		[]string{"b", "c"},
	}
	fmt.Println(p3)

	// 结构体内是一段连续的内存
	q := &pask{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &q.a) //0xc0000120da
	fmt.Printf("%p\n", &q.b) //0xc0000120db
	fmt.Printf("%p\n", &q.c) //0xc0000120dc

	// 结构体内存是对齐的
	q1 := &oask{
		a: int8(10),
		b: int8(20),
		c: "cC",
		d: "d",
	}
	fmt.Printf("%p\n", &q1.a) //0xc000064330
	fmt.Printf("%p\n", &q1.b) //0xc000064331
	fmt.Printf("%p\n", &q1.c) //0xc000064338  // 64位操作系统是跨8个字节,所以在第8个字节的时候使用内存对齐功能
	fmt.Printf("%p\n", &q1.d) //0xc000064348
}
