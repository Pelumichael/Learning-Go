package main

import "fmt"

// Slices are basically arrays without length -  unfixed sizes
// Writing Format

func main() {
	//[] <element_type>
	//var ids []string
	//var songNames []string
	var months []string // This is a Type of slice

	// Initialisation
	//songNames = []string{
	//	"Song1",
	//	"Song2",
	//	"Song3",
	//}
	months = []string{
		"Jan", // 0
		"Feb", // 1
		"Mar", // 2
		"Apr", // 3
		"May", // 4
		"Jun", // 5
		"Jul", // 6
		"Aug", // 7
		"Sep", // 8
		"Oct", // 9
		"Nov", // 10
		"Dec", // 11
	}
	//SLICING
	//<Slice or Array value>[<low_index>:<high_index>]
	allMonths := months[:12]
	halfYr := months[0:6]
	q1 := months[0:3]
	_q1 := months[0:4]
	anotherq1 := halfYr[:3]
	q2 := months[3:6]
	_q2 := halfYr[3:]
	q3 := months[6:9]
	q4 := months[9:12]
	december := months[11:12]
	fmt.Printf("allMonths = %v \n", allMonths)
	fmt.Printf("halfYr = %v \n", halfYr)
	fmt.Printf("q1 = %v \n", q1)
	fmt.Printf("_q1 = %v \n", _q1)
	fmt.Printf("q2 = %v \n", q2)
	fmt.Printf("q3 = %v \n", q3)
	fmt.Printf("q4 = %v \n", q4)
	fmt.Printf("december = %v \n", december)
	fmt.Printf("_q2 = %v \n", _q2)
	fmt.Printf("anotherq1 = %v \n", anotherq1)
	//fmt.Printf("_q2 with capacity = %v \n", cap(_q2))

	// SLICING WITH MAX ATTRIBUTE
	//<Slice or Array>[low_index:high_index:max]
	halfYr_2 := months[6:9:9]
	halfYr_3 := months[6:9]
	fmt.Printf("halfYr_2 = %v \n", cap(halfYr_2))
	fmt.Printf("halfYr_3 = %v \n", cap(halfYr_3))

	//Making a Slice(an empty slice)
	//_2023Months := []string{} //Literal method (not recommended)
	_2023Months := make([]string, 6) //Proper initialization of a slice
	fmt.Println(_2023Months)

	//Operations On Slice
	//Append
	fmt.Println(halfYr)
	halfYr_extended := append(halfYr, "monthX", "monthY")
	fmt.Println(halfYr_extended)

	//Copy
	_Q2 := make([]string, len(halfYr_extended))
	copy(_Q2, halfYr_extended)
	fmt.Println(_Q2)

	//Shallow copy
	items := []string{"brush", "phone", "chair"}
	_items := items
	fmt.Printf("address of items is: %[1]p\n", items)
	fmt.Printf("address of _items is: %[1]p\n", _items)

	_items = append(_items, "table")
	fmt.Println(items)
	fmt.Println(_items)
}
