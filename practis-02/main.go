package main

import (
	"fmt"
	"os"
)

/*
 学生管理系统
 写一个系统能够查看 新增 删除学生 函数版
*/

var (
	allStudent map[int]*student
)

type student struct {
	name string
	age  int
}

func newStudent(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}

func showAllStudent() {
	fmt.Println("All students are blow:")
	for key, value := range allStudent {
		fmt.Printf("id：%d, name: %s ,age: %d\n", key, value.name, value.age)
	}
	fmt.Println("========")
}

func addNewStudent() {
	var name string
	var age int
	var id int
	fmt.Printf("Please input student id: ")
	fmt.Scanln(&id)
	fmt.Printf("Please input student name: ")
	fmt.Scanln(&name)
	fmt.Printf("Please input student age: ")
	fmt.Scanln(&age)
	newGuy := newStudent(name, age)
	allStudent[id] = newGuy
}

func deleteStudent() {
	var id int
	fmt.Printf("Please input student id: ")
	fmt.Scanln(&id)
	delete(allStudent, id)
}

func main() {
	allStudent = make(map[int]*student, 50)
	for {
		fmt.Println("Welcome to tudent manager system, you can choose one of below in using numbers")
		fmt.Printf(` 
  1. show all student
  2. add new student
  3. delete student
  4. quit
Please enter your choice: `)
		var choiceNum int
		fmt.Scanln(&choiceNum)
		fmt.Printf("Your choice Num is %d \n", choiceNum)
		switch choiceNum {
		case 1:
			showAllStudent()
		case 2:
			addNewStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Your choice have been made is wrong, please choose again")
		}
	}

}
