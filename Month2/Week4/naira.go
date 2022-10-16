package main

import "fmt"

type naira float32

func (n naira) kobo() kobo {
	fmt.Printf("Converting %v Naira to Kobo  \n", n)
	return kobo(n * 100)
}

type kobo float32

func (k kobo) naira() naira {
	fmt.Printf("Converting %v kobo to Naira \n", k)
	return naira(k / 100)

}

func main() {
	convertKobo := kobo(10)
	fmt.Println(convertKobo.naira())

	convertNaira := naira(10)
	fmt.Println(convertNaira.kobo())
}
