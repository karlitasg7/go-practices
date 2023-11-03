package main

import (
	"fmt"

	"rsc.io/quote"
)

const Pi float32 = 3.1416

const (
	x = 100
	y = 0b1010 // binary
	z = 0o12   // octal
	w = 0xFF   // hexadecimal
)

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

	var firstName string = "Karla"
	var age int

	age = 18

	fmt.Println(firstName, age)

	var (
		lastName string = "Gonzalez"
		height   int    = 1
	)

	fmt.Println(lastName, height)

	var otherName, number = "Karla", 1

	fmt.Println(otherName, number)

	otherName2, number2 := "Karla", 1

	fmt.Println(otherName2, number2)

	fmt.Println(Pi)

	fmt.Println(x, y, z, w)
	fmt.Println(Friday)

}
