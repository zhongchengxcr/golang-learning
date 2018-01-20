package main

import "fmt"

func main() {

	var a1 [1]int
	var a2 [2]string

	a3 := [4]string{"z", "h", "o"}
	a4 := [10]int{4: 1}
	a5 := [...]int{2: 3, 1: 2, 0: 1}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)

	a5[2] = 10
	fmt.Println(a5)

	p := &a5

	fmt.Println(*p) // [1 2 10]
	fmt.Println(p)  //&[1 2 10] 指针数组

	x, y := 1, 2

	//在go中数据是值类型
	a6 := [...]*int{&x, &y}

	fmt.Println(a6)

	//长度为数据类型的一部分,不同类型不能比较和赋值
	a7 := [2]int{1, 2}
	a8 := [2]int{1, 2}
	fmt.Println(a7 == a8)

	//返回指向数组的指针 ,指针可以操作数组
	a9 := new([10]int)
	a9[2] = 1
	fmt.Println(*a9)

	a0 := [...][2]int{
		{1, 2},
		{3, 4},
		{5, 1: 6}}
	//[[0 0] [0 0] [0 0]]
	fmt.Println(a0)

	a10 := [...]int{9, 9, 3, 4, 5, 6, 7}

	fmt.Println(a10)

	fmt.Println("--------------------------")
	len := len(a10)


	//冒泡排序
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if a10[i] > a10[j] {
				temp := a10[i]
				a10[i] = a10[j]
				a10[j] = temp
			}
		}
	}

	fmt.Println(a10)

}
