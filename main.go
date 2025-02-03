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

	wg.Wait()
}

func generateId(wg *sync.WaitGroup, c chan string) {
	fmt.Println("Generando id...")
	id := uuid.New()
	sleepSeconds(3)
	fmt.Printf("El id generado es -> %s \n", id)
	sleepSeconds(1)
	fmt.Println("Enviando...")
	sleepSeconds(1)
	c <- id.String()
	wg.Done()
}

func receiveId(wg *sync.WaitGroup, idChan chan string) {
	id := <-idChan
	fmt.Println("Recibido")
	sleepSeconds(1)
	fmt.Printf("El id recibido es -> %s", id)
	wg.Done()
}

func sleepSeconds(number time.Duration) {
	time.Sleep(number * time.Second)
}
