package main

//TYPES

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hello Class 2.")

	//Declaring Variables
	//var aUint8 uint8
	//var aUint16 uint16
	//var aUint32 uint32
	//var aUint64 uint64
	//var aInt8 int8
	//var aInt16 int16
	//var aInt32 int32
	//var aInt64 int64

	//var a int
	//var a16 int8
	//var a32 int32
	////var aUint uint
	//
	//var accountNumber uint16
	var b string = "altschool"
	//var c float32
	//var d bool
	////var e int32
	////var f int64
	////
	////
	////aa := 30
	////bb := "altschool backend class"
	////cc := 3.145
	////[][][][][][][][] ==> 1 byte
	////[][][][][][][][] ==> 1 byte
	//
	////var complex64 complex64
	////var complex128 complex128
	//
	//fmt.Printf("a = %v [%T, %d bits] \n", a, a, unsafe.Sizeof(a)*8)
	//fmt.Printf("a16 = %v [%T, %d bits] \n", a16, a16, unsafe.Sizeof(a16)*8)
	//fmt.Printf("a32 = %v [%T, %d bits] \n", a32, a32, unsafe.Sizeof(a32)*8)
	//fmt.Printf("accountNumber = %v [%T, %d bits] \n", accountNumber, accountNumber, unsafe.Sizeof(accountNumber)*8)
	//fmt.Printf("d = %v [%T, %d bits] \n", d, d, unsafe.Sizeof(d)*8)
	//fmt.Printf("c = %v [%T, %d bits] \n", c, c, unsafe.Sizeof(c)*8)
	fmt.Printf("b = %v [%T, %d bits] \n", b, b, unsafe.Sizeof(b)*8)
}
