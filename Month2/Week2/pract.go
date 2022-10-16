package main

import "fmt"

func main() {
	num := []int{0, 1, 2, 3, 4, 5, 6}
	for i, val := range num {
		fmt.Printf("THe value at index %v is %v \n", i, val)
	}

}

//Write a function that gives the value of an index within a range using for loops.
