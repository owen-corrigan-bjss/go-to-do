package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	toDos "to-do-app"
	"to-do-app/helpers"
)

func PrintJsonToDos(toDos ...toDos.ToDoJson) string {
	toDoStr := "Here are your todo's:\n"
	for i, v := range toDos {
		toDoStr += fmt.Sprintf("%d: %s\n", i+1, v.Description)
	}
	fmt.Println(toDoStr)
	return toDoStr
}

func PrintToDosList(toDos ...toDos.ToDo) string {
	toDoStr := "Here are your todo's:\n"
	for i, v := range toDos {
		toDoStr += fmt.Sprintf("%d: %s\n", i+1, v.Description)
	}
	fmt.Println(toDoStr)
	return toDoStr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	breakLoop := false

	for !breakLoop {
		toDoJson, err := os.Open("/Users/owen.corrigan/projects/go-to-do/toDos.json")
		if err != nil {
			log.Fatal(err)
		}

		defer toDoJson.Close()

		fmt.Printf("what would you like to do:\n1: print to do from JSON list\n2: print toDos from an array\n0: exit\n")

		scanner.Scan()
		input := scanner.Text()

		if input == "0" {
			breakLoop = true
		} else if input == "1" {
			decodedJson := helpers.DecodeJson(toDoJson)
			PrintJsonToDos(decodedJson...)
		} else if input == "2" {

			PrintToDosList(toDos.ToDoList...)
		}
	}
}
