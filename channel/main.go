package main

import "fmt"


type Vihecle interface{
	Drive()
}

type Car struct {
	Make string
}

func (c Car) Drive() {
	fmt.Println("Driving a car:", c.Make)
}

type Bike struct {
	Make string
}

func (b Bike) Drive() {
	fmt.Println("Driving a car:", c.Make)
}

func main() {

	myCar := Car{Make: "Toyota"}
	
	var v Vihecle 	= myCar
}