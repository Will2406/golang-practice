package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}

	fmt.Println("Comenzar los trabajos...")
	go Worker("Daniel", "Hacer la puerta", wg)
	wg.Add(1)
	go Worker("Belen", "Hacer el baño", wg)
	wg.Add(1)
	go Worker("Yanella", "Hacer la kk", wg)
	wg.Add(1)

	wg.Wait()

	fmt.Println("Se termino el trabajo")

}

func Worker(name string, work string, wg *sync.WaitGroup) {

	fmt.Printf("El trabajador %s comenzo a %s \n", name, work)
	for i := 0; i < 10; i++ {
		fmt.Printf("El trabajo '%s' el proceso es %d/10\n", work, i+1)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}
