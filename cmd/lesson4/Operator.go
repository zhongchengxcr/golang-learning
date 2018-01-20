package main

import "fmt"

/**

运算符 从左至右结合

优先级  高 -> 低

 -	^ , !     																	( 一元运算符)
 -	* , / , % , << , >> , & , &^ , + , - , | , ^ , == , != , < , <= , >= , >    ( 二元运算符)
 -  <-
 -	&&
 -	||

*/

func main() {
	fmt.Println(2 ^ 1)
	fmt.Println(!true)  //false
	fmt.Println(2 << 3) //16
	fmt.Println(8 >> 3) //1
	fmt.Println("--------------------------")
	// 4个 位运算符   & , | , ^ , &^
	//		 6: 0110
	//		10: 1010
	//--------------
	// &	    0010 =2      两个都为1 的时候 是 1
	// |		1110 =14    只要有一个是1 的时候  是 1
	// ^		1100 =12    只有一个是 1 的时候  是 1
	// &^		0100 =4     第二个数是一 的情况下 吧第一个数对应的位上 变成 0

	fmt.Printf("%b\n", 6&10)
	fmt.Printf("%b\n", 6|10)
	fmt.Printf("%b\n", 6^10)
	fmt.Printf("%b\n", 6&^10)

	fmt.Println(6 & 10)
	fmt.Println(6 | 10)
	fmt.Println(6 ^ 10)
	fmt.Println(6 &^ 10)
}
