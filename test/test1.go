package main

import (
	"fmt"
)

// Define the Engine struct
type Engine struct {
	Horsepower int
 }
 // Define the Car struct that embeds Engine
 type Car struct {
	Brand  string
	Model  string
	Year   int
	Engine Engine
}

// Method to start the engine
func (e *Engine) Start() {
	fmt.Printf("Engine with %d horsepower is starting...\n", e.Horsepower)
 }
 // Method to drive the car
 func (c *Car) Drive() {
	fmt.Printf("Driving a %d %s %s...\n", c.Year, c.Brand, c.Model)
	c.Engine.Start()
 }
 func main() {
	car := Car {
	   Brand:  "Toyota",Model:  "MR2",Year:   1991,
	   Engine: Engine{Horsepower: 332},
	}
	car.Drive()
 }
 