package main

import (
	"fmt"
	"sync"
)

var (
	balance = 100
)

func deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	balance += amount
	lock.Unlock()
}

func getBalance() int {
	a := balance
	return a
}

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go deposit(i*100, &wg, &lock)
	}

	wg.Wait()
	fmt.Println(getBalance())
}
