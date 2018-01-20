package main

import "fmt"

type Student struct {
	Name string
}

//私有公有 是相对于包而言的
type Teacher struct {
	Name string
}

func (s *Student) say(prefix string) {
	s.Name = "kyn"
	fmt.Println(prefix, s.Name)
}

func (t Teacher) say(prefix string) {
	t.Name = "kyn"
	fmt.Println(prefix, t.Name)
}

type Tc int

/**
只能绑定在本包内
*/
func (tc *Tc) Print() {
	fmt.Println(*tc)
}

func (tc Tc) Printx() {
	fmt.Println(tc)
}

type Inc int

func (i *Inc) inc(num int) {
	*i = *i + Inc(num)
}

func main() {

	s := Student{Name: "zc"}
	s.say("Hello ")

	t := Teacher{Name: "zc"}
	t.say("Hi ")

	fmt.Println(s)
	fmt.Println(t)

	var tc Tc
	tc = 9

	tc.Print()
	(*Tc).Print(&tc)

	Tc.Printx(tc)

	fmt.Println("====================")

	var a Inc
	b := Inc(1)
	fmt.Println(b)

	fmt.Println(a)
	a.inc(100)
	fmt.Println(a)

}
