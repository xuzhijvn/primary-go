/**
类型查询:
	类型查询，就是根据变量，查询这个变量的类型
*/
package main

import (
	"fmt"
)

func main() {
	// 定义一个interface{}类型变量，并使用string类型值”abc“初始化
	var a interface{} = "abc"

	// 在switch中使用 变量名.(type) 查询变量是由哪个类型数据赋值。
	switch v := a.(type) {
	case string:
		fmt.Println("字符串")
	case int:
		fmt.Println("整型")
	default:
		fmt.Println("其他类型", v)
	}

	var value1, ok1 = a.(string)
	fmt.Println(value1)
	fmt.Println(ok1)
}
