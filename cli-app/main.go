package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	types "github.com/owen-corrigan-bjss/to-do-app/to-do-types"
)

var fp = fmt.Printf

func CreateToDo(list *types.ToDoListContainer, ids *types.IdCounter, scanner *bufio.Scanner) {
	fmt.Println("\nwhat task would you like to add:")
	scanner.Scan()
	input := scanner.Text()

	list.CreateToDoItem(input, ids)
	fp("\n%s\n", list.ListToDos())

}

func UpdateToDoStatus(list *types.ToDoListContainer, ids *types.IdCounter, scanner *bufio.Scanner) {
	fmt.Println("\nwhich task do you want to update?:")
	scanner.Scan()
	input := scanner.Text()
	_, err := list.UpdateToDoItemStatus(input)
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fp("\n%s\n", list.ListToDos())
	}
}

func DeleteToDo(list *types.ToDoListContainer, ids *types.IdCounter, scanner *bufio.Scanner) {
	fmt.Println("\nwhich task do you want to delete?:")
	scanner.Scan()
	input := scanner.Text()
	_, err := list.DeleteToDoItem(input)

	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fp("\n%s\n", list.ListToDos())
	}
}

func main() {

	toDoList := types.NewToDoList()
	toDoIds := types.NewCounter()

	scanner := bufio.NewScanner(os.Stdin)

	breakLoop := false

	for !breakLoop {

		fmt.Println("what would you like to do:\n1: list to do's\n2: create a new to do\n3: edit to do status\n4: delete to do\n0: exit")

		scanner.Scan()
		input := scanner.Text()

		command := strings.Split(input, " ")[0]

		if command == "0" {
			breakLoop = true
		} else if command == "1" {
			fp("%s\n", toDoList.ListToDos())
		} else if command == "2" {
			CreateToDo(toDoList, toDoIds, scanner)
		} else if command == "3" {
			UpdateToDoStatus(toDoList, toDoIds, scanner)
		} else if command == "4" {
			DeleteToDo(toDoList, toDoIds, scanner)
		}
	}
}
