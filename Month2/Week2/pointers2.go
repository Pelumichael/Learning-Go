package main

import "fmt"

func main() {
	pointer()
}
func pointer() {
	// 1 2 3 4 5 6 7 8 9
	//[][][][][][][][][] 0
	//[][][][][][][][][] 1
	//[][][][][][][][][] 2
	//[][][4][][][][][][] 3
	//[][][][][][][][][] 4
	//[][][][][][][][][] 5
	//[][][][][][][][][] 6
	//[][][][][][][][][] 7
	//[][][][][][][][][] 8

	var four int = 4
	var addressOfFour *int = &four

	//four => identifier
	//addressOfFour => pointer

	fmt.Printf("The address of Four is: %v \n", addressOfFour)
	fmt.Printf("The Value at the address of four is: %v \n", *addressOfFour)
	// Updating the address by calling the value.
	four = 43
	fmt.Printf("The Value at the address of four is now: %v \n", four)
	fmt.Printf("The address of Four is: %v \n", addressOfFour)

	//	Updating the address by calling the address.
	*addressOfFour = 50
	fmt.Printf("value of four is %v \n", four)

	structPtr := &struct{ x, y int }{2, 4}
	fmt.Printf("struct=%#v, type=%T\n", structPtr, structPtr)
}
