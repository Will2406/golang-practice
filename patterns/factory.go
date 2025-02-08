package main

import "fmt"

type IProduct interface {
	getStock() int
	setStock(stock int)
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

type Desktop struct {
	Computer
}

func makeLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 23,
		},
	}
}

func makeDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 45,
		},
	}
}

func GetComputerFactory(name string) (IProduct, error) {
	if name == "Desktop" {
		return makeDesktop(), nil
	}

	if name == "Laptop" {
		return makeLaptop(), nil
	}

	return nil, fmt.Errorf("Invalid computer type")
}

/*func main() {
	newComputerType := "Desktop"
	computer, _ := GetComputerFactory(newComputerType)

	fmt.Println(computer.getName())
}*/
