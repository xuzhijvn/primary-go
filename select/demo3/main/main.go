package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 3)

	var i1 = "one"

	for {
		select {
		case i1 = <-c1:
			fmt.Println("received ", i1)
		case c1 <- i1:
			fmt.Println("sent ", i1)
		default:
			fmt.Printf("no communication\n")
		}
		time.Sleep(1 * time.Second)
	}
}
