package main

import (
	"fmt"
	"os"
)

var TODOLIST = []string{"Default Todo"}
var menuNumber int

func main() {
	welcome()
	menu()
}

func newLine(numberOfLines int) {
	i := 0
	for i < numberOfLines {
		fmt.Print("\n")
		i++
	}
}

func welcome() {
	fmt.Println("********{Welcome to the TODO CLI app}*******")
	newLine(1)
}

func listTodo() {
	fmt.Println("List Of Items in Your Todo list are:")
	newLine(
		1)
	for index, item := range TODOLIST {
		fmt.Printf("TODO #%v: %v", index+1, item)
	}
	newLine(2)
	menu()
}

/*func displayMenu() {
	fmt.Println("Select Operation")
	fmt.Println("1. Create Todo \t\t 2. List Todo items")
	fmt.Println("3. Edit Todo items\t 4. Delete Todo items")
	fmt.Println("0. Exit Program \t")
	newLine(2)
}*/

func createTodo() {
	fmt.Println("Creating Todo Items")
	menu()
}

func editTodo() {
	fmt.Println("Editing todo item")
	menu()
}

func deleteTodo() {
	fmt.Println("Deleting todo item")
	menu()
}

func exitProgram() {
	fmt.Println("Goodbye.")
	os.Exit(0)
}

func menu() {
	newLine(1)
	fmt.Println("Select Operation")
	fmt.Println("1. Create Todo \t\t 2. List Todo items")
	fmt.Println("3. Edit Todo items\t 4. Delete Todo items")
	fmt.Println("0. Exit Program ")
	newLine(2)

	_, err := fmt.Scan(&menuNumber)
	if err != nil {
		fmt.Println("Error: please select the correct menu Item.")
	}

	switch menuNumber {

	case 1:
		createTodo() //fmt.Println("Create todo Items")
	case 2:
		listTodo() //fmt.Println("List the todo Items")
	case 3:
		editTodo() //fmt.Println("Edit todo Items")
	case 4:
		deleteTodo() //fmt.Println("Delete  todo  Items")
	case 0:
		exitProgram()
	default:
		fmt.Println("Error: Invalid Input")
	}
}
