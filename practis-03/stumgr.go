package main

import "fmt"

type student struct {
	name string
	age  int
	id   int
}

type stuManger struct {
	allStudents map[int]*student
}

func (s stuManger) showAllStudents() {
	for _, val := range s.allStudents {
		fmt.Printf("学号: %d  姓名: %s 年龄: %d \n", val.id, val.name, val.age)
	}
}

func (s stuManger) addNewStudent() {
	var (
		name string
		age  int
		id   int
	)
	fmt.Print("请输入学号: ")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名: ")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄: ")
	fmt.Scanln(&age)
	newStu := &student{
		name: name,
		age:  age,
		id:   id,
	}
	s.allStudents[id] = newStu
}

func (s stuManger) updateStduent() {
	var (
		id   int
		name string
	)
	fmt.Print("请输入修改的学号:")
	fmt.Scanln(&id)
	stu, err := s.allStudents[id]
	if !err {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("学号: %d  姓名: %s 年龄: %d \n", stu.id, stu.name, stu.age)
	fmt.Print("请输入修改的姓名:")
	fmt.Scanln(&name)
	stu.name = name
	fmt.Println("修改完毕，修改后的信息如下:")
	fmt.Printf("学号: %d  姓名: %s 年龄: %d \n", stu.id, stu.name, stu.age)
}

func (s stuManger) deleteStudent() {
	var (
		id int
	)
	fmt.Print("请输入删除的学号:")
	fmt.Scanln(&id)
	_, err := s.allStudents[id]
	if !err {
		fmt.Println("查无此人")
		return
	}
	delete(s.allStudents, id)
}
