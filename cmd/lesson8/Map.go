package main

import (
	"fmt"
	"sort"
)

func main() {

	//map 查找比普通线性查找快,但是比数组和slice慢100倍
	var m1 = map[int]string{1: "go", 2: "scala"}
	m2 := make(map[int]string)

	m2[3] = "elrong"
	delete(m1, 1)

	fmt.Println(m1)
	fmt.Println(m2)

	m3 := map[int]map[int]string{}
	m4 := make(map[int]map[int]string)

	m4[2] = map[int]string{1: "java"}
	v1, ok := m4[2]

	fmt.Println("====================================================================================")

	v2 := &v1
	fmt.Println(v2, ok)

	fmt.Println(m3)

	fmt.Println(m4[1])

	var a1 = []string{"java", "go", "scala", "python"}
	fmt.Println(a1)
	//此处v是a1值得copy,对v的操作不会改变a1
	for i, v := range a1 {
		fmt.Println(i, v)
		v = "ok"
	}
	fmt.Println(a1)

	fmt.Println("========================")
	//此处v是a1值得copy,对v的操作不会改变a1
	var m5 = map[string]string{"one": "zc", "two": "kyn", "three": "pop"}
	fmt.Println(m5)
	for k, v := range m5 {
		fmt.Println(k, v)
		v = "ok"
	}
	fmt.Println(m5)

	fmt.Println("================sort map ================")

	sm := map[int]string{1: "java", 2: "go", 3: "scala", 4: "python"}
	sa := make([]int, len(sm))
	i := 0
	for k, _ := range sm {
		sa[i] = k
		i++
		fmt.Println(k)
	}

	sort.Ints(sa)

	for _, v := range sa {
		v := sm[v]
		fmt.Println(v)
	}

	fmt.Println("================exchange ================")
	sm1 := map[int]string{1: "java", 2: "go", 3: "scala", 4: "python"}

	sm2 := map[string]int{}
	for k, v := range sm1 {
		sm2[v] = k
	}



	fmt.Println(sm1)
	fmt.Println(sm2)
}
