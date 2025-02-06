package main

import (
	"fmt"
	"time"
)

func main() {
	y := func() int {
		return 2 * 2
	}()

	fmt.Println(y)
	c := make(chan int)
	go func() {
		fmt.Println("start function..")
		time.Sleep(5 * time.Second)
		fmt.Println("end")
		c <- 1
	}()
	<-c
}
