package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println()
	fmt.Println("1. 添加学员信息")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 退出系统")
	fmt.Println()
}

func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按照要求输入学员信息")
	fmt.Print("请输入学员的学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员的班级：")
	fmt.Scanf("%s\n", &class)
	return newStudent(id, name, class)
}

func main() {
	sm := newStudentMgr()
	for {
		showMenu()

		var input int
		fmt.Print("请输入你要操作的序号：")
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			stu := getInput()
			sm.addStudent(stu)
			stuInfo := stu.stuInfo()
			fmt.Printf("%s，添加成功\n", stuInfo)
		case 2:
			stu := getInput()
			ok := sm.modifyStudent(stu)
			if ok {
				stuInfo := stu.stuInfo()
				fmt.Printf("%s，编辑成功\n", stuInfo)
			}
		case 3:
			sm.showStudent()
		case 4:
			fmt.Println("系统退出")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
