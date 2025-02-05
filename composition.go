package main

import "fmt"

type Vehicle struct {
	id    int
	speed int
}

func (v *Vehicle) Start() {
	fmt.Println("starting.. ")
}

type Motorcycle struct {
	Vehicle
	seats int
	size  int
}

// composition over inheritance
// https://medium.com/deep-golang/why-golang-choose-composition-as-its-base-rather-than-inheritance-1225d22a4798
