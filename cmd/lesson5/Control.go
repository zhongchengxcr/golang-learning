package main

import (
	"fmt"
	"strconv"
)

const (
	B int64 = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {
	A := 9
	var p *int = &A
	fmt.Println(*p)

	a := 1
	a++

	fmt.Println(a)

	if a++; a < 4 {
		fmt.Println("a:" + strconv.Itoa(a))
	}

	//无限循环
	for {
		fmt.Println("zhongc")
		break
	}

	b := 1

	// 相当于 while(吧<5)
	for b < 5 {
		b++
		fmt.Println(b)
	}

	fmt.Println()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	c := 1
	switch c {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

	switch d := 1; {
	case d >= 0:
		fmt.Println("c >= 0")
		fallthrough
	case d >= 1:
		fmt.Println("c >= 1")
	case d >= 2:
		fmt.Println("c >= 2")
	default:
		fmt.Println("default")
	}

	//默认break只能跳出1层循环
LABLE1:
	for {
		for i := 1; i < 3; i++ {

			fmt.Println(i)
			if i == 2 {
				break LABLE1
				//continue LABLE1 继续执行标签所在循环
			}

			if i == 1 {
				goto LABLE2
			}
		}
	}
LABLE2:
}
