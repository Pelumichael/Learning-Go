package main

import (
	"fmt"
	"os"
)

var TODOLIST = []string{"Default Todo"}

func main() {
	welcome()
	displayMenu()
	var menuNumber int
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

func newLine(numberOfLines int) {
	i := 0
	for i < numberOfLines {
		fmt.Print("\n")
		i++
	}
}

func welcome() {
	fmt.Println("********{Welcome to the TODO CLI app}*******")
	newLine(3)
}

func listTodo() {
	fmt.Println("List Of Items in Your Todo list are:")
	newLine(
		1)
	for index, item := range TODOLIST {
		fmt.Printf("TODO #%v: %v", index+1, item)
	}
	newLine(2)
}

func displayMenu() {
	fmt.Println("Select Operation")
	fmt.Println("1. Create Todo \t")
	fmt.Println("2. List Todo \t")
	fmt.Println("3. Edit Todo \t")
	fmt.Println("4. Delete Todo \t")
	fmt.Println("0. Exit Program \t")
	newLine(2)
}

func createTodo() {
	fmt.Println("Creating Todo Items")
}

func editTodo() {
	fmt.Println("Editing todo item")
}

func deleteTodo() {
	fmt.Println("Deleting todo item")
}

func exitProgram() {
	fmt.Println("Goodbye.")
	os.Exit(0)
}
