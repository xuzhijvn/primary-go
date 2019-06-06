package reflect

import (
	"fmt"
	"reflect"
)

type T struct{}

func main() {
	t := &T{}

	res := t.compute(11, 11, "Add")

	fmt.Println(res)

}

func (t *T) compute(a int, b int, method string) int64 {
	x := reflect.ValueOf(a)
	y := reflect.ValueOf(b)
	in := []reflect.Value{x, y}
	res := reflect.ValueOf(t).MethodByName(method).Call(in)[0].Int()
	return res
}

//方法名首字母需要大写
func (t *T) Add(a int, b int) int {
	return a + b
}

//方法名首字母需要大写
func (t *T) Multi(a int, b int) int {
	return a * b
}
