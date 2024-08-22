package types

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type ToDo struct {
	description string
	completed   bool
}

type ToDoList map[string]ToDo

type ToDoListContainer struct {
	list ToDoList
	lock sync.Mutex
}

func NewToDoList() *ToDoListContainer {
	list := ToDoListContainer{}
	list.list = make(map[string]ToDo)
	return &list
}

func (list *ToDoListContainer) ListToDos() string {

	toDoListToPrint := "here are your to do's:\n"

	if len(list.list) == 0 {
		toDoListToPrint = "\nthere are currently no items in your list\n"
	}
	for i, v := range list.list {
		toDoListToPrint += fmt.Sprintf("id: %s desc: %s : complete: %t\n", i, v.description, v.completed)
	}
	return toDoListToPrint
}

func (list *ToDoListContainer) CreateToDoItem(description string, counter *IdCounter) {
	list.lock.Lock()

	key := counter.GetNewId()

	list.list[key] = ToDo{description, false}

	list.lock.Unlock()
}

func (list *ToDoListContainer) UpdateToDoItemStatus(id string) error {
	list.lock.Lock()

	itemToUpdate, ok := list.list[id]

	if !ok {
		return errors.New("item doesn't exist")
	}

	itemToUpdate.completed = !itemToUpdate.completed

	list.list[id] = itemToUpdate

	list.lock.Unlock()
	return nil
}

func (list *ToDoListContainer) DeleteToDoItem(id string) error {
	list.lock.Lock()

	_, ok := list.list[id]

	if !ok {
		return errors.New("item doesn't exist")
	}

	delete(list.list, id)

	list.lock.Unlock()
	return nil
}

type IdCounter struct {
	count int
	lock  sync.Mutex
}

func (count *IdCounter) IncrementCounter() {
	count.count++
}

func (count *IdCounter) GetNewId() string {
	count.lock.Lock()
	defer count.lock.Unlock()
	id := count.count
	count.IncrementCounter()
	return strconv.Itoa(id)
}

func NewCounter() *IdCounter {
	ids := IdCounter{}
	ids.count = 0
	return &ids
}
