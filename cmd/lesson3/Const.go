package main

import "fmt"

// 常量的值在编译的时候就已经确定了,所以在运行的时候无法改变
const x int = 1
const y = 'A'
const v1, v2 = 1, 2
const (
	v3 = 1
	v4 = 2
	v5
)

const (
	a = 'A'
	b = "A"
	c = iota
	d = iota
	e
)

const (
	f = iota
)

func main() {
	fmt.Println(v3) //1
	fmt.Println(v4) //2
	fmt.Println(v5) //2

	fmt.Println()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e) //4
	fmt.Println(f) //0


	
}
