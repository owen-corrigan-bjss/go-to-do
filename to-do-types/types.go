package types

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type ToDo struct {
	Description string
	Completed   bool
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
		toDoListToPrint += fmt.Sprintf("id: %s desc: %s : complete: %t\n", i, v.Description, v.Completed)
	}
	return toDoListToPrint
}

func (list *ToDoListContainer) GetSingleToDo(id string) ToDo {
	return list.list[id]
}

func (list *ToDoListContainer) GetToDoMap() ToDoList {
	return list.list
}

func (list *ToDoListContainer) CreateToDoItem(description string, counter *IdCounter) string {
	list.lock.Lock()
	defer list.lock.Unlock()

	key := counter.GetNewId()

	list.list[key] = ToDo{description, false}

	return key

}

func (list *ToDoListContainer) UpdateToDoItemStatus(id string) (toDo ToDo, err error) {
	// list.lock.Lock()
	// defer list.lock.Unlock()

	itemToUpdate, ok := list.list[id]

	if !ok {
		return ToDo{}, errors.New("item doesn't exist")
	}

	itemToUpdate.Completed = !itemToUpdate.Completed

	list.list[id] = itemToUpdate

	return itemToUpdate, nil
}

func (list *ToDoListContainer) DeleteToDoItem(id string) (deleted bool, err error) {
	list.lock.Lock()
	defer list.lock.Unlock()

	_, ok := list.list[id]

	if !ok {
		return false, errors.New("item doesn't exist")
	}

	delete(list.list, id)

	return true, nil
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
