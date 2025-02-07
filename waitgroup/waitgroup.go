package main

import "fmt"

func main() {
	fmt.Println("Comenzar los trabajos...")
	Worker("Daniel", "Hacer la puerta")
	Worker("Belen", "Hacer el baño")
	Worker("Yanella", "Hacer la kk")
}

func Worker(name string, work string) {

	fmt.Printf("El trabajador %s comenzo a %s \n", name, work)
	for i := 0; i < 10; i++ {
		fmt.Printf("El trabajo '%s' el proceso es %d/10\n", work, i+1)
	}
}
