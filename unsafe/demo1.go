package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {

	fmt.Println(&x)
	fmt.Println(unsafe.Pointer(&x))
	fmt.Println(uintptr(unsafe.Pointer(&x)))

	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"

	//不安全
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb1 := (*int16)(unsafe.Pointer(tmp))
	*pb1 = 109

	fmt.Println(x.b) // "42"

}
