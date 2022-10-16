package main

import "fmt"

type altschoolMember bool

type address struct {
	street      string
	city, state string
	houseNumber uint8
}

type diameter int

type planetName struct {
	long  string
	short string
}

type person struct {
	name        string
	age         uint8
	address     address
	email       address
	gender      string
	phone       uint16
	altschooler altschoolMember
}
type planet struct {
	diameter
	planetName
	desc string
}

type currency struct {
	name    string
	country string
	code    string
}

func main() {
	//DECLARING NAME STRUCTS

	//type car struct {
	//	make         string
	//	model        string
	//	year         uint8
	//	color        string
	//	mileage      uint8
	//	transmission string
	//	owners       []person //If multiple persons driver the car.
	//	status       bool
	//}

	//Initialising structs

	//var firstPerson = new(person)
	//
	//firstPerson.name = "John Doe"
	//firstPerson.gender = "Male"
	//firstPerson.age = 43
	//firstPerson.phone = 000000000
	//firstPerson.address = address{
	//	city:        "Lagos",
	//	state:       "Lagos State",
	//	street:      "Wesely Street",
	//	houseNumber: 72,
	//}
	//
	//fmt.Println(firstPerson)

	//var usd = currency{"United States Dollar", "United States", "usd"}
	//fmt.Println(usd)

	//ANNONYMOS FIELDS

	earth := planet{
		diameter: 7000,
		planetName: planetName{
			long:  "Earth",
			short: "E",
		},
		desc: "Earth is the third planet",
	}

	//fmt.Println(earth)

	//PROMOTED FIELDS
	mars := planet{}
	mars.diameter = 3000
	mars.desc = "It is the 4th Planet"
	mars.long = "Mars"
	mars.short = "M"

	//STRUCT AS PARAMETERS
	goToPlanet(&earth)
	fmt.Printf("We are still in Planet %s", earth.long)
}

func goToPlanet(p *planet) {
	p.long = "Jupiter"
	fmt.Printf("Here we going to Planet %s \n", p.long)
}
