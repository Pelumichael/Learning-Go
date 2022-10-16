package main

import "fmt"

type automobile uint

// method or Receiver Functions
func (a automobile) start1() uint {
	//fmt.Println("Starting start1 Automobile")
	return uint(a * 4)
}

// Function
func start2() {
	fmt.Println("Starting start2 Automobile")
}

func main() {
	start2()

	toyota := automobile(1)
	fmt.Println(toyota.start1())
}
