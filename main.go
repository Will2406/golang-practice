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

	wg.Add(2)
	go generateId(wg, c)
	go receiveId(wg, c)

	car := Car{
		id:    1,
		name:  "name",
		brand: false,
	}

	car2 := NewCar(2, "default", true)
	fmt.Println(car)
	fmt.Println(*car2)

	motorcycle := Motorcycle{}
	motorcycle.id = 1
	motorcycle.speed = 78
	motorcycle.Start()

	dog := Dog{name: "bobby"}
	person1 := Person{name: "Daniel", age: 12}

	GetMessage(dog)
	GetMessage(person1)

	wg.Wait()
}

func generateId(wg *sync.WaitGroup, c chan<- string) {
	for i := 0; i < 2; i++ {
		id := uuid.New()
		c <- fmt.Sprintf("%d. %s orginal \n", i, id)
	}
	close(c)
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
