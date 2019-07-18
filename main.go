package main

import (
	"fmt"
	"os"
)

/*
sims -- SStudent Information Management System

1. 添加学员信息
2. 编辑学员信息
3. 展示学员信息
4. 退出系统

*/

func showMenu() {
	fmt.Println("----------------------")
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 退出系统")
	fmt.Println("----------------------")

}

//获取用户输入的学员信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("----------------------")
	fmt.Println("请按要求输入学员信息")
	fmt.Println("请输入学员的学号：")
	_, _ = fmt.Scanf("%d\n", &id)
	fmt.Println("请输入学员的姓名：")
	_, _ = fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学员的班级：")
	_, _ = fmt.Scanf("%s\n", &class)

	stu := newStudent(id, name, class)
	return stu
}

func main() {
	sm := newStudentMgr()
	for {
		//1. 打印系统菜单
		showMenu()
		//2. 等待用户选择要执行的选项
		var input int
		fmt.Println("请输入选项(1-4):")
		_, _ = fmt.Scanf("%d\n", &input)
		//3. 执行用户选择的动作
		switch input {
		case 1:
			//添加学员
			newStu := getInput()
			sm.addStudent(newStu)
		case 2:
			//编辑信息
			newStu := getInput()
			sm.editStudent(newStu)
		case 3:
			//列出学员信息
			sm.allStudent()
		case 4:
			//退出系统
			os.Exit(0)
		default:
			fmt.Println("输入的选项有误，请重新输入")
		}
	}

}
