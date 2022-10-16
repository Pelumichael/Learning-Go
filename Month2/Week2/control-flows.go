package main

import "fmt"

func main() {
	//x, y := 12, 23
	//if x > y {
	//	fmt.Println("X is greater than Y")
	//} else {
	//	fmt.Println("X is lesser than Y")
	//}
	//
	////	While and do...while loop
	//for y < 10{...}
	//
	//// Infinite loop
	//for {
	//
	//}
	// Tradition loop
	//for c := 0; c < 10; c++ { //++ means +1 i.e +=1
	//	//3 conditions
	//	// Start the loop
	//	//	Check the status of the loop
	//	//	Keep the loop running
	//	fmt.Printf("C is %v \n", c)
	//}

	// For Range (Slices, Array,Stream,Maps and channels)
	//	for i
	//	val := range []string{}

	daysOfTheWeek := []string{"Monday", "Tuesday", "Wednesday"}

	// If statements
	// switch statements

	// FOR LOOPS
	// break, continue, goto
	// Infinite
	// For LOOPS
	//Four(4) variants of For loops

	//var i int = 0
	////Infinite
	//for {
	//	if i < len(daysOfTheWeek) {
	//		newVar := daysOfTheWeek[i]
	//		i++
	//		fmt.Printf("The value at the index %v is %v \n", i, newVar)
	//		continue
	//	} else {
	//		fmt.Printf("Sorry i is not less than 3")
	//		break
	//	}
	//}

	//	TRADITIONAL
	//for i := 0; i < len(daysOfTheWeek); i++ {
	//	val := daysOfTheWeek[i]
	//	fmt.Printf("The value at the index %v is %v \n", i, val)
	//}

	//For Ranges of Composite Types (Slices, Array,Stream,Maps and channels)

	//for i, val := range daysOfTheWeek {
	//	fmt.Printf("The value at index %v is %v \n", i, val)
	//}

	// For Conditions
	var s int
	for s < len(daysOfTheWeek) {
		newVar := daysOfTheWeek[s]
		fmt.Printf("The value at index %v is %v \n", s, newVar)
		s++
	}

}
