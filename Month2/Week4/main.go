package main

import "fmt"

type engine struct {
	fuel   int
	thrust int
	hp     int
}

func (e *engine) start() {
	fmt.Print("Engine has Started \n")
}

type automobile struct {
	make  string
	model string
	engine
	electric bool
}

func (a *automobile) drive() {
	fmt.Printf("Automobile %s %s is on the road \n", a.make, a.model)
}

type plane struct {
	automobile
	engineCount int
	fixedWings  bool
	maxAltitude int
}

func (p *plane) fly() {
	fmt.Printf("Aircraft %s %s is clear for Takeoff! \n", p.make, p.model)
}

func (p *plane) land() {
	fmt.Printf("Aircraft %s %s is landing Now! \n", p.make, p.model)
}
func main() {
	benz := &automobile{
		make:  "Mercedes Benz",
		model: "c300",
		engine: engine{
			fuel:   30,
			thrust: 4000,
			hp:     400,
		},
		electric: false,
	}

	benz.start()
	benz.drive()

	audi := &automobile{}
	audi.make = "BMW"
	audi.model = "X5 series"
	audi.fuel = 45
	audi.hp = 600
	audi.thrust = 400
	audi.electric = true

	audi.start()
	audi.drive()

	plane := &plane{
		automobile: automobile{
			make:  "Boeing",
			model: "737", engine: engine{fuel: 300000000,
				thrust: 40000,
				hp:     50000},
			electric: false,
		},
		engineCount: 2,
		fixedWings:  true,
		maxAltitude: 1209904,
	}
	plane.start()
	plane.fly()
}
