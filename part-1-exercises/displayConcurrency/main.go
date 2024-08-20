package main

import (
	"fmt"
	"sync"
	toDos "to-do-app/part-1-exercises"
)

var fp = fmt.Printf

type lockedToDo struct {
	toDos []toDos.ToDo
	lock  sync.Mutex
}

var list []toDos.ToDo = []toDos.ToDo{
	{Description: "here is task 1", Complete: false},
	{Description: "here is task 2", Complete: true},
	{Description: "here is task 3", Complete: false},
	{Description: "here is task 4", Complete: false},
	{Description: "here is task 5", Complete: true},
	{Description: "here is task 6", Complete: true},
	{Description: "here is task 7", Complete: false},
	{Description: "here is task 8", Complete: true},
	{Description: "here is task 9", Complete: false},
	{Description: "here is task 10", Complete: true},
}

func ReadToDoDesc(toDoStruct *lockedToDo, index int, done chan bool) {
	fp("task %d desc: %s\n", index, toDoStruct.toDos[index].Description)
	done <- true
}

func ReadToDoStatus(toDoStruct *lockedToDo, index int, done chan bool) {
	fp("task %d status: %v\n", index, toDoStruct.toDos[index].Complete)
	done <- true
}

func main() {
	var toDoList lockedToDo
	toDoList.toDos = list
	done := make(chan bool)

	for i := 0; i < 10; i++ {

		toDoList.lock.Lock()
		go ReadToDoDesc(&toDoList, i, done)
		<-done
		go ReadToDoStatus(&toDoList, i, done)
		<-done
		toDoList.lock.Unlock()
	}
}
