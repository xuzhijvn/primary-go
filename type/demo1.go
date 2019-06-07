/**
类型定义：
	类型定义可以在原类型的基础上创造出新的类型，有些场合下可以使代码更加简洁
*/
package main

import (
	"fmt"
)

// 定义一个需要五个参数的函数类型
type handle func(str string, str2 string, num int, money float64, flag bool)

// exec函数，接收handle类型的参数
func exec(f handle) {
	f("hello", "world", 10, 11.23, true)
}

func method1(str string, str2 string, num int, money float64, flag bool) {
	fmt.Println("method1: ", str, str2, num, money, flag)
}
func main() {

	exec(method1)

	// 定义一个函数类型变量，这个函数接收五个字符串类型的参数
	var method2 = func(str string, str2 string, num int, money float64, flag bool) {
		fmt.Println("method2: ", str+str2, num+100, money-1, flag)
	}
	exec(method2)

	// 匿名函数作为参数直接传递给exec函数
	exec(func(str string, str2 string, num int, money float64, flag bool) {
		fmt.Println("method3: ", str)
	})

}
