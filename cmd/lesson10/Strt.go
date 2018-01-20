package main

import (
	"go-learing/cmd/lesson10/model"
	"fmt"
)

type Person struct {
	Id   int
	Name string
	//匿名结构
	Son struct {
		Name string
		age  int
	}
}

//匿名字段
type Nms struct {
	int
	string
}

func main() {

	s1 := model.Student{}
	s1.Class = "one"
	s1.Name = "zc"
	s1.Age = 18
	fmt.Println(s1)

	s2 := model.Student{
		Name:  "kyn",
		Age:   10,
		Class: "two",
	}

	fmt.Println(s2)
	A(&s2)
	fmt.Println(s2)

	s3 := &model.Student{
		Name:  "zdy",
		Age:   16,
		Class: "three",
	}

	fmt.Println(s3)
	A(s3)
	fmt.Println(s3)

	fmt.Println("=============================================")

	nm := &struct {
		Id   int
		Name string
	}{
		Id:   1,
		Name: "zc",
	}

	fmt.Println(nm)

	fmt.Println("=============================================")

	s4 := model.Student{
		Name:  "kyn",
		Age:   10,
		Class: "two",
	}

	s5 := model.Student{}

	s5 = s4

	fmt.Println(s4)
	fmt.Println(s5)

	fmt.Printf("%p \n", &s4) // 0xc42007c330
	fmt.Printf("%p \n", &s5) //0xc42007c360

	fmt.Println("=============================================")

	s7 := &model.Student{
		Name:  "kyn",
		Age:   10,
		Class: "two",
	}

	var s6 *model.Student
	s6 = s7

	fmt.Println(s6)
	fmt.Println(s7)

	fmt.Printf("%p \n", s6) // 0xc42007c360
	fmt.Printf("%p \n", s7) //0xc42007c360

	fmt.Println("=============================================")

	s8 := &model.Student{
		Name:  "kyn",
		Age:   10,
		Class: "two",
	}

	s9 := &model.Student{
		Name:  "kyn",
		Age:   10,
		Class: "two",
	}

	var s10 *model.Student

	s10 = s9

	fmt.Println(s8)
	fmt.Println(s9)

	fmt.Println(*s8 == *s9) //true
	fmt.Println(s8 == s9)   //false

	fmt.Println(s10 == s9)

	fmt.Println("=============================================")

	var t1 Teacher
	var t2 Teacher = Teacher{
		age:   10,
		name:  "kyn",
		human: human{Sex: 0},
	}

	t1.human.age = 19
	t1.age = 18
	t1.name = "zc"
	t1.Sex = 1

	fmt.Println(t1)
	fmt.Println(t2)
}

type human struct {
	Sex int8
	age int
}

type Teacher struct {
	human
	age  int
	name string
}

//struct 的传递是 值传递,若想引用传递,可以使用指针
func A(stu *model.Student) {
	stu.Age = 30
	fmt.Println(stu)
}
