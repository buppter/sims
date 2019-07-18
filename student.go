package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

//学员管理的类型
type studentMgr struct {
	allStudents []*student
}

//newStudent 是student类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id,
		name,
		class,
	}
}

//newStudentMgr 是studentMgr的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

//添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
	fmt.Printf("学生[学号：%d，姓名：%s，班级：%s]添加成功!\n", newStu.id, newStu.name, newStu.class)
}

//编辑学生信息
func (s *studentMgr) editStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id {
			s.allStudents[i] = newStu
			fmt.Println("学生信息修改成功！")
			return
		}
	}

	fmt.Printf("输入的学生信息有误，系统中不存在学号是：%d 的学生 \n", newStu.id)
}

//展示学生信息
func (s *studentMgr) allStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d，姓名：%s，班级：%s \n", v.id, v.name, v.class)
	}
}
