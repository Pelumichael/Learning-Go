package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type TODO struct {
	Id   string
	Todo string
}

var TODOLIST []TODO

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
	newLine(2)
	for index, item := range TODOLIST {
		fmt.Printf("TODO #%v: %v %v", index+1, item.Id, item.Todo)
		newLine(2)
	}

	menu()
}

func createTodo() {
	fmt.Println("Please Enter your Todo Item")
	var todoInput string
	_, err := fmt.Scan(&todoInput)
	if err != nil {
		fmt.Println("Error: Please enter a valid todo item")
	}

	var todo TODO
	todo.Id = strconv.Itoa(rand.Intn(10000000000))
	todo.Todo = todoInput

	TODOLIST = append(TODOLIST, todo)
	fmt.Println("Todo Item Created.")
	menu()
}

func updateTodo() {
	var updateItem string
	var updateInput string
	fmt.Println("Enter Id of the Item You want to Update: ")
	_, err := fmt.Scan(&updateItem)
	if err != nil {
		fmt.Println("Error: Please select the correct menu Item")
	}
	for index, item := range TODOLIST {
		if item.Id == updateItem {
			TODOLIST = append(TODOLIST[:index], TODOLIST[index+1:]...)
			var newTodo TODO
			//Todo: display current todo item to user
			fmt.Println("Enter your todo: ")
			_, err := fmt.Scan(&updateInput)
			if err != nil {
				fmt.Println("Error: Please select the correct menu Item")
			}
			newTodo.Id = updateItem
			newTodo.Todo = updateInput

			TODOLIST = append(TODOLIST, newTodo)

			fmt.Println("Todo item updated")
			menu()
		}
	}
	fmt.Println("Todo item not Found.")
	menu()
}

func deleteTodo() {
	var deleteItem string
	fmt.Println("Enter the ID of the Todo Item to Delete. ")
	_, err := fmt.Scan(&deleteItem)
	if err != nil {
		fmt.Println("Error: Please select the correct menu Item")
	}

	for index, item := range TODOLIST {
		if item.Id == deleteItem {
			TODOLIST = append(TODOLIST[:index], TODOLIST[index+1:]...)
			fmt.Println("Todo item deleted")
			menu()
		}
	}
	fmt.Println("TODO item not found")
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

		menu()
	}

	switch menuNumber {

	case 1:
		createTodo() //fmt.Println("Create todo Items")
	case 2:
		listTodo() //fmt.Println("List the todo Items")
	case 3:
		updateTodo() //fmt.Println("Edit todo Items")
	case 4:
		deleteTodo() //fmt.Println("Delete  todo  Items")
	case 0:
		exitProgram()
	default:
		fmt.Println("Error: Invalid Input")
	}
}
