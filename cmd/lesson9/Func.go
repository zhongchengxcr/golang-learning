package main

import "fmt"

func main() {

	fmt.Println("===================func A()=====================")
	a, b := A(1, 2)
	fmt.Println(a, b)
	fmt.Println("===================func B()=====================")

	var c = 3
	var d = 4
	var e = 5

	//s := make([]int, 3)

	B(c, d, e)

	//fmt.Println(s)

	fmt.Println(c, d, e)

	fmt.Println("===================func C()=====================")

	cc := []int{1, 2, 3}

	//当传递为数组类型时,函数内的改变,不会对原始值进行改变,因为数组是值类型,函数内得到的是拷贝
	//当传递的参数是 slice 时,函数内的操作,会对原始值进行改变,因为slice是 引用类型, 函数内得到的是 slice的引用(指针)
	//如果想在函数内操作外部的 值类型,可以使用指针,见 func D()
	C(cc)
	fmt.Println(cc)

	fmt.Println("===================func D()=====================")

	dd := [3]int{1, 2, 3}

	D(&dd)
	fmt.Println(dd)

	fmt.Println("===================func E()=====================")

	ee := E
	ee()

	fmt.Println("===================func F()=====================")

	ff := func() {
		fmt.Println("hello func without name!")
	}

	ff()

	fmt.Println("===================func 闭包=====================")

	clofunc := Closure(1)

	clo1 := clofunc(2)
	clo2 := clofunc(3)

	fmt.Println(clo1)
	fmt.Println(clo2)

	fmt.Println("=================== defer =====================")

	// defer fmt.Println("A")
	// defer fmt.Println("B")

	// for i := 0; i < 3; i++ {
	// 	defer fmt.Println(i)
	// }

	// for i := 0; i < 3; i++ {
	// 	//return 才会执行,即函数主题运行完毕的时候
	// 	fmt.Println("%p",&i)
	// 	defer func (){
	// 		//输出333,因为defer配合匿名函数在使用的时候 , 在for循环执行之后defer才开始执行,此时i为3,
	// 		i=i+1
	// 		fmt.Println("defer  %p",&i)
	// 		fmt.Println(i)
	// 	}()
	// }

	fmt.Println("=================== panic/recover =====================")

	

	P1()
	P2()
	P3()



}

func P1() {
	fmt.Println("func P1")
}

func P2() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in P2")
		}
	}()

	panic("panic in p2")
	fmt.Println("func P2")
}

func P3() {
	fmt.Println("func P3")
}

func Closure(x int) func(int) int {
	fmt.Println("%p\n", &x)
	return func(y int) int {
		//此处的x为x的地址,并不是值得copy
		fmt.Println("%p\n", &x)
		return x + y
	}
}

func E() {
	fmt.Println("hell go")
}

func D(dd *[3]int) {

	dd[0] = 1
	dd[1] = 1
	dd[2] = 1
}

func A(a, b int) (ra int, rb string) {
	ra = 3
	rb = "go"
	return ra, rb
}

func B(a ...int) string {
	fmt.Println(a)

	a[0] = 9
	a[1] = 9
	a[2] = 9

	return "java"
}

func C(c []int) {
	fmt.Println(c)
	c[0] = 9
	c[1] = 9
	c[2] = 9

}

func mainx() {

	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		//i作为参数传递给 fmt.Println函数 , 因为i是int类型,所以 为值传递
		defer fmt.Println("defer i = ", i)
		defer func() {
			//这个里面的i为内存地址
			fmt.Println("defe_closure i = ", i)
		}()

		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}

	for _, f := range fs {
		f()
	}

	//closure i = 4
	//closure i = 4
	//closure i = 4
	//closure i = 4
	//defe_closure i =4
	//defer i = 3
	//defe_closure i =4
	//defer i = 2
	//defe_closure i =4
	//defer i = 1
	//defe_closure i =4
	//defer i = 0
	
}