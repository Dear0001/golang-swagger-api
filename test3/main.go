package main

import "fmt"

type Speaker interface {
	Speak() string
}
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

	
type Cat struct{}

func (c Cat) Speak() string {
   return "Meow!"
}

func main() {
   animals := []Speaker{
	Dog{}, Cat{},
}
   for _, animal := range animals {
      fmt.Println(animal.Speak())
   }
}