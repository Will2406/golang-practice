package main

type Car struct {
	id    int
	name  string
	brand bool
}

func NewCar(id int, name string, brand bool) *Car {
	return &Car{
		id:    id,
		name:  name,
		brand: brand,
	}
}

// Best practices for the creation of new instances
// https://go.dev/doc/effective_go#allocation_new
// https://stackoverflow.com/questions/18125625/constructors-in-go
