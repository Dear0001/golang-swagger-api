package main

import "fmt"

type Address struct {
	City, State string
 }
 
 type Person struct {
	Name    string
	Address Address
 }
 func main() {
	p := Person{
	   Name:    "Alice",
	   Address: Address{City: "Wonderland", State: "Dreams"},
	}
	fmt.Printf("%s lives in %s, %s\n", p.Name, p.Address.City, p.Address.State)
 }
  