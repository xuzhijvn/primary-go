/**
类型别名：
	1、类型别名是给原类型取了一个小名，本质上没有发生任何变化。
	2、类型别名，只能对包内的类型产生作用。
*/
package main

import (
	"fmt"
)

// 根据string类型，定义类型S
type S string

func (r *S) Hi() {
	fmt.Println("S hi")
}

// 定义S的类型别名为T
type T = S

func (r *T) Hello() {
	fmt.Println("T hello")
}

// 函数参数接收S类型的指针变量
func execute(obj *S) {
	obj.Hello()
	obj.Hi()
}

func main() {
	t := new(T)
	s := new(S)
	execute(s)
	// 将T类型指针变量传递给S类型指针变量
	execute(t)
}
