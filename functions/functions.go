package main

import (
	"fmt"
	"time"
)

func suma(values ...int) int {
	a := 0
	for _, num := range values {
		a += num
	}
	return a
}

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

	sumTotal := suma(1, 2, 3, 4, 5, 6)
	fmt.Println(sumTotal)
	<-c
}
