package main

import "fmt"

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"DZD", "Algerian Dollar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"EUR", "Euro", "Italy", 978},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"EUR", "Euro", "Greece", 978},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"HKD", "Hong Kong Dollar", "Hong Kong", 344},
	Curr{"USD", "US Dollar", "United States", 840},
}

func isDollar(curr Curr) bool {
	var result bool

	switch curr {
	default:
		result = false
	case Curr{"AUD", "Australian Dollar", "Australia", 36}:
		result = true
	case Curr{"HKD", "Hong Kong Dollar", "Hong Kong", 344}:
		result = true
	case Curr{"USD", "US Dollar", "United States", 840}:
		result = true
	}
	return result
}

func isDollar2(curr Curr) bool {
	dollars := []Curr{currencies[1], currencies[6], currencies[7]}
	switch curr {
	default:
		return false
	case dollars[0]:
		fallthrough
	case dollars[1]:
		fallthrough
	case dollars[2]:
		return true
	}
}

func isEuro(curr Curr) bool {
	switch curr {
	case currencies[2], currencies[4], currencies[10]:
		return true
	default:
		return false
	}
}

func main() {
	curr := Curr{"EUR", "Euro", "Italy", 978}
	if isDollar(curr) {
		fmt.Printf("%+v is Dollar currency \n", curr)
	} else if isEuro(curr) {
		fmt.Printf("%+v is Euro currency \n", curr)
	} else {
		fmt.Println("Currency is not dollar or Euro")
	}

	dol := Curr{"HKD", "Hong Kong Dollar", "Hong Kong", 344}
	if isDollar2(dol) {
		fmt.Println("Dollar currency found: ", dol)
	}

	fmt.Println(len(currencies))
}
