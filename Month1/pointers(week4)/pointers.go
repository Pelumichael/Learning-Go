package main

import "fmt"

////POINTERS
//[1][][][][][][][]
//[][][3][][][][][]
//[][][][][][][][]
//[][4][][][][][][]
//[][][][][][][][]

func main() {
	//Declaration and Initialization
	var _ uint = 12
	var name string = "pointer"
	var namePtr *string = &name
	fmt.Println(*namePtr)
	//Declaration
	//var agePtr *uint = &age
	//fmt.Println(agePtr)
	//0xc00000e098 = agePtr
	//0xc00000e0d0 = noBirthday
	//Your birthday is: 12
	//0xc00000e098 = Birthday
	//Your birthday is 13

	//noBirthday(age)
	//fmt.Printf("Your birthday is: %v ", age)
	//fmt.Println("\n")
	//birthday(&age)
	//fmt.Printf("Your birthday is %v", age)

	//TERMINAL OUTPUT
	//Your birthday is: 12
	//
	//Your birthday is 13

	//fmt.Printf("The address of variable age is %v", agePtr)
}

//func birthday(x *uint) {
//	fmt.Println(x)
//fmt.Println(*x) = 12(i.e value of x)
//*x = *x + 1
//Your birthday is: 12
//
//0xc00000e098
//Your birthday is 13

//}

//func noBirthday(x uint) {
//	fmt.Println(&x)
//x = x * 1
//0xc00000e0b0 = for noBirthday
//Your birthday is: 12
//
//0xc00000e098 = for Birthday
//Your birthday is 13
//}

// * ==> value of
// & ==>  address of
