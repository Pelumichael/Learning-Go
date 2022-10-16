package main

import "fmt"

// Simple Conversion System
type ounce float64

func (o ounce) cup() cup {
	return cup(o * 0.125)
}

type cup float64

func (c cup) quart() quart {
	return quart(c * 0.25)
}

func (c cup) ounce() ounce {
	return ounce(c * 8.0)
}

type quart float64

func (q quart) gallon() gallon {
	return gallon(q * 0.25)
}

func (q quart) cup() cup {
	return cup(q * 4)
}

type gallon float64

func (g gallon) quart() quart {
	return quart(g * 4)
}

func main() {
	quart1 := quart(1) //Quart is an Object
	quart2 := quart(2)

	fmt.Println(quart1)
	fmt.Println(quart2)

	//gal := gallon(5)
	//fmt.Printf("There are %v quarts in %v gallons", gal.quart(), gal)
}
