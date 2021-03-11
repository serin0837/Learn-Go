package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string // slice of string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
}
