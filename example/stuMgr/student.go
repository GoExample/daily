package main

import "fmt"

type student struct {
	Id    int
	Name  string
	Class string
}

func newStudent(id int, name, class string) *student {
	return &student{
		Id:    id,
		Name:  name,
		Class: class,
	}
}

type studentMgr struct {
	allStudents []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 1000),
	}
}

func (s *student) stuInfo() string {
	return fmt.Sprintf("学号：%d 姓名：%s 班级：%s", s.Id, s.Name, s.Class)
}

func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

func (s *studentMgr) modifyStudent(newStu *student) bool{
	for i, v := range s.allStudents {
		if v.Id == newStu.Id {
			s.allStudents[i] = newStu
			return true
		}
	}
	fmt.Printf("输入的学生信息有误，系统中没有学号是：%d的学生\n", newStu.Id)
	return false
}

func (s *studentMgr) showStudent() {
	if len(s.allStudents) == 0 {
		fmt.Println("系统中无任务学员信息")
		return
	}

	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.Id, v.Name, v.Class)
	}
}
