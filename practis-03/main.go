package main

import (
	"fmt"
	"os"
)

var manager stuManger

// 结构体版的学生管理系统
func msgSupport() {
	println(`
	欢迎使用学生管理系统
 
	1. 查看所有学生信息
	2. 新增学生信息
	3. 修改学生信息
	4. 删除学生信息
	5. 退出
	`)
}

func main() {
	manager = stuManger{
		allStudents: make(map[int]*student, 50),
	}
	for {
		msgSupport()
		fmt.Print("请选择: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			manager.showAllStudents()
		case 2:
			manager.addNewStudent()
		case 3:
			manager.updateStduent()
		case 4:
			manager.deleteStudent()
		case 5:
			os.Exit(0)
		}

	}
}
