package main

import "fmt"

/*
Array(A collection of values of the same type)
*/
// How to instantiate an Array
//var ages [100]int
//var days [7]string
//var weekdays [5]string
//var truth [256]bool
//var data [4][2]int         // Multiple dimensional Array
//var datA [5]map[string]int //A compositeType: Joining 2 or types together, its type is always on the right side.

// Format
// [Length]<element_type>
type days [7]string

func main() {


	// Initialization
	var days1 *days = new(days)
	days1[0] = "Sunday"
	days1[1] = "Monday"
	days1[2] = "Tuesday"
	days1[3] = "Wednesday"
	days1[4] = "Thursday"
	days1[5] = "Friday"
	days1[6] = "Saturday"
	fmt.Print(days1)
	printDaysOfTheWeek(days1)
	fmt.Printf("\n")
	fmt.Println(days1)

	//fmt.Print(days)
	// Result: [Monday Tuesday Wednesday Thursday Friday Saturday Sunday]

	//var board = [5][2]int{
	//	{12, 3},
	//	{12, 4},
	//	{12, 1},
	//	{1, 1},
	//	{1, 3},
	//}
	//
	//var shortForms = [2][2][2][2]int{
	//	{{{4, 4}, {3, 2}, {3, 5}, {2, 3}}},
	//	{{{4, 4}, {3, 2}, {3, 5}, {2, 3}}},
	//}



	// For transversing through an array.
	//for key, val := range days {
	//	fmt.Printf("For keys=> %v, the value is: %v \n", key, val)
	//	//Result
	//For keys=> 0, the value is: Monday
	//For keys=> 1, the value is: Tuesday
	//For keys=> 2, the value is: Wednesday
	//For keys=> 3, the value is: Thursday
	//For keys=> 4, the value is: Friday
	//For keys=> 5, the value is: Saturday
	//For keys=> 6, the value is: Sunday
}

// Passing Arrays as Parameters
func printDaysOfTheWeek(daysArr *days) {
	for key, _ := range daysArr {
		daysArr[key] = "Weekend"
		//fmt.Printf("For key=> %v, the value is: %v \n", key, val)
	}
}
