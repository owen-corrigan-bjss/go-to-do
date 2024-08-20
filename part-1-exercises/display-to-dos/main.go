package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	toDos "to-do-app/part-1-exercises"
	"to-do-app/part-1-exercises/helpers"
)

func PrintJsonToDos(toDos ...toDos.ToDoJson) string {
	toDoStr := "Here are your todo's:\n"
	for i, v := range toDos {
		toDoStr += fmt.Sprintf("%d: %s\n", i+1, v.Description)
	}
	fmt.Println(toDoStr)
	return toDoStr
}

func OpenAndPrintJson() {
	toDoJson, err := os.Open("/Users/owen.corrigan/projects/go-to-do/toDos.json")
	if err != nil {
		log.Fatal(err)
	}
	defer toDoJson.Close()
	decodedJson := helpers.DecodeJson(toDoJson)
	PrintJsonToDos(decodedJson...)
}

func CreateAndPrintJson() {
	newToDoJson, err := os.Open("/Users/owen.corrigan/projects/go-to-do/part-1-exercises/toDos.json")
	if err != nil {
		log.Fatal(err)
	}
	defer newToDoJson.Close()
	CreateNewJsonFile()
	decodedJson := helpers.DecodeJson(newToDoJson)
	PrintJsonToDos(decodedJson...)
}

func PrintToDosList(toDos ...toDos.ToDo) string {
	toDoStr := "Here are your todo's:\n"
	for i, v := range toDos {
		toDoStr += fmt.Sprintf("%d: %s\n", i+1, v.Description)
	}
	fmt.Println(toDoStr)
	return toDoStr
}

func CreateNewJsonFile() {
	toDosForJson := []toDos.ToDo{{Description: "this is a new JSON file", Complete: false}}

	for i := 1; i < 10; i++ {
		newToDo := toDos.ToDo{Description: fmt.Sprintf("this is new task %d", i+1), Complete: false}
		toDosForJson = append(toDosForJson, newToDo)
	}

	file, err := os.Create("/Users/owen.corrigan/projects/go-to-do/display-to-dos/newToDos.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	toDoList := toDos.ToDoListForConvert{}

	toDoList.ToDos = toDosForJson[:]

	jsonFormattedToDos, err := json.Marshal(toDoList)

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(jsonFormattedToDos)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	breakLoop := false

	for !breakLoop {

		fmt.Printf("what would you like to do:\n1: print to do from JSON list\n2: print toDos from an array\n3: create a new JSON file and print it\n0: exit\n")

		scanner.Scan()
		input := scanner.Text()

		if input == "0" {
			breakLoop = true
		} else if input == "1" {
			OpenAndPrintJson()
		} else if input == "2" {
			PrintToDosList(toDos.ToDoList...)
		} else if input == "3" {
			CreateAndPrintJson()
		}
	}
}
