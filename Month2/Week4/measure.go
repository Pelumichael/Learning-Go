package main

import "fmt"

type millimeter float32
type centimeters float32
type meter float32
type kilometer float32

func (c centimeters) millimeter() float32 {
	return float32(c * 10)
}

func (c centimeters) meter() float32 {
	return float32(c / 100)
}

func (m millimeter) centimeters() centimeters {
	fmt.Printf("converting %v millimeter to centimeter \n", m)
	return centimeters(m / 10)
}

func (m millimeter) meter() meter {
	fmt.Printf("converting %v millimeter to meter \n", m)
	return meter(m.centimeters() / 100)
}

func main() {
	length := centimeters(100)
	fmt.Println(length)
	fmt.Println(length.millimeter())

	length2 := millimeter(2000)
	fmt.Println(length2.meter())

	//length1 := millimeter(2)
	//length2 := millimeter(90)
	//
	//fmt.Println(length1.centimeters().meter())
	//fmt.Println(length2.centimeters())
}
