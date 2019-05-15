package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	var i1, i2 string
	i2 = "two"

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("get ", <-c2)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c3 <- "three"
	}()

	select {
	case i1 = <-c1:
		fmt.Println("received ", i1)
	case c2 <- i2:
		fmt.Println("sent ", i2)
		//延迟一秒再退出程序，让程序有时间打印出“get two”
		time.Sleep(1 * time.Second)
	case i3, ok := <-c3: // same as: i3, ok := (<-c3)
		if ok {
			fmt.Println("received ", i3)
		} else {
			fmt.Println("c3 is closed")
		}
		//default:
		//	fmt.Printf("no communication\n")
	}
}
