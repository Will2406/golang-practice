package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) getInfo() string {
	return fmt.Sprintf("Mi nombre es %s y mi edad %d", p.name, p.age)
}

type Dog struct {
	name string
}

func (d Dog) getInfo() string {
	return fmt.Sprintf("Soy un perro y mi nombre es %s", d.name)
}

type PrintInfo interface {
	getInfo() string
}

func GetMessage(p PrintInfo) {
	fmt.Println(p.getInfo())
}
