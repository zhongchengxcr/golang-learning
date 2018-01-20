package main

//可以使用别名
import io "fmt"

const PI = 3.14

const (
	const1 = "const"
	const2 = "const"
)

var (
	str1 = "str"
	str2 = "str"
)

var (
	k1 string
	k2 string
)

var p1, p2, p3 = 1, 2, 3

var age = 2

type name string


//小写代表是private ,不可以被外部调用
var pop = "pop"

//可见性规则
//大写代表是public ,可以被外部调用   规则试用 常量,变量,类型,接口,结构,函数
var Sex = "girl"

type student struct {
}

type golang interface {
}

func main() {
	var p4, p6, p9 = 1, 2, 3
	var (
		v5=6
		v8=0
	)
	io.Print(p4)
	io.Print(p6)
	io.Print(p9)
	io.Print(v5)
	io.Print(v8)

	io.Println("Hello  main " + pop)
}
