package main

import (
	"fmt"
	"strconv"
)

// 长度1个字节 ,不可以用数字代替
var flag bool = true

//根据平台决定是多少位 32/64 , 2的 32/64 次方
var num int = 1

//-127 ~ 128
var num8 int8 = 123

// 0 ~ 255 , TODO 为什么可以超出取值范围 ans :编译会报错
var num9 uint8 = 12

// uint8 的 别名
var num0 byte = 32

var num32 int32 = 8888

// rune 是 int32 的别名
var num33 rune = 999

//go 中没有double类型  , float32 站位:32/8 = 4个字节 精确到7位小数 ,
var decimal32 float32 = 3.32

// float 精确到 15位小数
var decimal64 float64 = 9.63876238746

//复数
var fu complex64
var fu1 complex128

/**

其他值类型
	- array struct string

引用类型
	- slice map  chan

接口 interface
函数 func

*/

/**


类型零值 , 值类型默认都是0,布尔是false , string为字符串

引用类型 为 nil ,相当于java中的null
*/

var v1 string  // ""
var v2 int     // 0
var v3 bool    // false
var v4 float64 // 0
var v5 []int   //空数组 这个是切片    []
var v6 [4]int  //空数据  这个是数组     [0 0 0 0]
var v7 []bool  // []
var v8 [3]bool //[false false false]

func main() {
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)
	fmt.Println(v7)
	fmt.Println(v8)

	type (
		byte     int8
		文本       string
		ByteSize int64 //增加代码可读性
	)

	var v9 文本 = "牛逼"

	fmt.Println(v9)

	var a1 float32 = 1.1

	// Go 不存在隐士转换 , 所有的转换都必须显示的声明
	//转换只能发生在相互兼容的类型之间

	var a2 int = int(a1)

	fmt.Println(a2)

	var a3 = 65

	a4 := string(a3)
	a5 := strconv.Itoa(a3)

	fmt.Println(a4) //  A
	fmt.Println(a5) //  65

}
