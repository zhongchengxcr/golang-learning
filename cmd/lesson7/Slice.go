package main

import "fmt"

//Slice
func main() {

	var s1 []int
	fmt.Println(s1) //[]

	var a1 = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(a1)
	//返回的是 slice 数组改变也会被改变
	s1 = a1[2:]
	s3 := a1[:2]
	s2 := a1[2:len(a1)]

	//[0 0 0]  2,3,4 包涵开始索引,不包含结束
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("============")
	// param 类型 ,长度(会初始化0值) ,容量  ,当超出容量的时候回翻倍容量
	s4 := make([]int, 3, 10)

	fmt.Println(s4)

	fmt.Println(len(s4), cap(s4))

	fmt.Println("============")

	a2 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	fmt.Println(a2)
	fmt.Println(len(a2), cap(a2))
	s5 := a2[2:5]
	s6 := a2[:2]

	fmt.Println("===== s5=======")
	fmt.Println(s5)
	fmt.Println(len(s5), cap(s5))

	fmt.Println("===== s6=======")
	fmt.Println(s6)
	fmt.Println(len(s6), cap(s6))
	s7 := s5[:6]

	fmt.Println("===== s7=======")
	fmt.Println(s7)

	fmt.Println(len(s7), cap(s7))

	//ASCII数组
	pop := []byte(`A`)

	fmt.Println(pop)

	s8 := make([]int, 3, 6)

	var s9 *[]int = &s8

	s8 = append(s8, 3, 4)

	fmt.Println(s8)
	fmt.Println(s9)

	fmt.Println("========================================")
	t1 := []int{1, 2, 3, 4, 5}
	t2 := t1[2:5]
	t3 := t1[1:3]

	fmt.Println(t2, t3)

	t1[2] = 9
	fmt.Println(t2, t3)

	fmt.Println("===================copy=====================")

	o1 := []int{1, 2, 3, 4, 5, 6}
	o2 := []int{7, 8, 9}

	copy(o1, o2)
	fmt.Println(o1, o2)

	o3 := []int{1, 2, 3, 4, 5, 6}
	o4 := []int{7, 8, 9}

	copy(o4, o3)
	copy(o4[:2], o3[2:5])

	o3[0] = 99
	fmt.Println(o3, o4)
	fmt.Println("========================================")


	q1 := []int{1,2,3,4,5,6}
	q2 := q1[:]

	fmt.Println(q1,q2)


}
