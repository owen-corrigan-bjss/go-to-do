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

func ReadToDoDesc(toDoStruct *lockedToDo, index int, wg *sync.WaitGroup) {
	fp("task %d desc: %s\n", index, toDoStruct.toDos[index].Description)
	wg.Done()
}

func ReadToDoStatus(toDoStruct *lockedToDo, index int, wg *sync.WaitGroup) {
	fp("task %d status: %v\n", index, toDoStruct.toDos[index].Complete)
	wg.Done()
}

func main() {
	var toDoList lockedToDo
	toDoList.toDos = list
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(2)
		toDoList.lock.Lock()
		go ReadToDoDesc(&toDoList, i, &wg)
		go ReadToDoStatus(&toDoList, i, &wg)
		toDoList.lock.Unlock()
	}

	wg.Wait()
}
