package main


import (
	"fmt"
)

type Animal struct {
	Name string
}
	
func (a *Animal) SetName(name string) {
	a.Name = name
}

func main() {
	a := Animal{Name: "Ah ki"}
	fmt.Println(a.Name) // Outputs: Buddy
	a.SetName("Ah do")
	fmt.Println(a.Name) // Outputs: Max
}