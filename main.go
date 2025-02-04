package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

func main() {
	wg := &sync.WaitGroup{}
	c := make(chan string)

	wg.Add(3)
	go generateId(wg, c)
	go receiveId(wg, c)
	go generateIdFake(wg, c)

	wg.Wait()
}

func generateId(wg *sync.WaitGroup, c chan<- string) {
	for i := 0; i < 100; i++ {
		id := uuid.New()
		sleepSeconds(2)
		c <- fmt.Sprintf("%d. %s orginal \n", i, id)
	}
	wg.Done()
}

func generateIdFake(wg *sync.WaitGroup, c chan<- string) {
	for i := 0; i < 45; i++ {
		id := uuid.New()
		sleepSeconds(1)

		c <- fmt.Sprintf("%d. %s fake \n", i, id)
	}

	wg.Done()
}

func receiveId(wg *sync.WaitGroup, idsChan <-chan string) {
	for id := range idsChan {
		fmt.Println(id)
	}
	wg.Done()
}

func sleepSeconds(number time.Duration) {
	time.Sleep(number * time.Second)
}
