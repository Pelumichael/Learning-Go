package main

import (
	"fmt"
)

// func <identifier> ([<argument list>]) (result list) {
// return
// }

var global int = 1000

func main() {
	prime := 37
	fmt.Printf("before double() %v \n", &prime)
	_ = double(&prime)
	fmt.Printf("after double() %v \n", prime)

	//fmt.Println(global)

	//fmt.Println(doubledValue)

	//boolResponse, intResponse := isPrime(prime)
	//fmt.Println(boolResponse, intResponse)
	//fmt.Println(primeResponse)
	//fmt.Printf("Is %v a prime number?  %v It is a Prime Number. \n", prime, isPrime(prime))

	// Assigning a Function To a variable
	//var isPrimeHAndler func(int) (bool, int) = isPrime
	//
	//fmt.Println(isPrimeHAndler(54))
}

// A Function that doubles whatever is assigned to it.
func double(amount *int) int {
	//fmt.Printf("The global variable is %v", global)
	//fmt.Println(amount)
	*amount = *amount * 2
	return *amount
}

//func play() {
//	fmt.Println("I am Playing.")
//}
//
//// Function to check if a number is Prime
//func isPrime(num int) (bool, int) {
//	lim := int(math.Sqrt(float64(num)))
//	for p := 2; p <= lim; p++ {
//		if (num % p) == 0 {
//			return false, num
//		}
//	}
//
//	return true, num
//}
