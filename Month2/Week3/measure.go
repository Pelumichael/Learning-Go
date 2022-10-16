package main

import "fmt"

type centimeters float32
type meter float32

func (c centimeters) millimeter() float32 {
	return float32(c * 10)
}

func (c centimeters) meter() float32 {
	return float32(c / 100)
}

type millimeter float32

func (m millimeter) centimeters() centimeters {
	fmt.Printf("converting %v millimeter to centimter \n", m)
	return centimeters(m / 10)
}

func main() {
	length1 := millimeter(2)
	length2 := millimeter(90)

	fmt.Println(length1.centimeters().meter())
	fmt.Println(length2.centimeters())
}
