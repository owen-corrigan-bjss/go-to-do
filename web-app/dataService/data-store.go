package dataService

import (
	types "github.com/owen-corrigan-bjss/to-do-app/to-do-types"
)

type ToDo struct {
	Description string
	Completed   bool
}

type DataStore interface {
	ListToDos() string
	GetSingleToDo(id string) (toDo ToDo, err error)
	GetToDoMap() ToDoList
	CreateToDoItem(description string, counter *types.IdCounter) string
	UpdateToDoItemStatus(id string) (toDo ToDo, err error)
	DeleteToDoItem(id string) (deleted bool, err error)
}

type DataStoreStruct struct {
	inMemoryToDoList *types.ToDoListContainer
	ids              *types.IdCounter
}

func NewDataStore() *DataStoreStruct {
	ds := DataStoreStruct{}
	ds.inMemoryToDoList = types.NewToDoList()
	ds.ids = types.NewCounter()
	return &ds
}

type ToDoList map[string]ToDo

func (ds *DataStoreStruct) ListToDos() string {
	return ds.inMemoryToDoList.ListToDos()
}

func (ds *DataStoreStruct) GetSingleToDo(id string) (toDo types.ToDo, err error) {
	todo, err := ds.inMemoryToDoList.GetSingleToDo(id)
	return todo, err
}

func (ds *DataStoreStruct) GetToDoMap() types.ToDoList {
	return ds.inMemoryToDoList.GetToDoMap()
}

func (ds *DataStoreStruct) CreateToDoItem(description string) string {
	return ds.inMemoryToDoList.CreateToDoItem(description, ds.ids)
}

func (ds *DataStoreStruct) UpdateToDoItemStatus(id string) (toDo types.ToDo, err error) {
	return ds.inMemoryToDoList.UpdateToDoItemStatus(id)
}

func (ds *DataStoreStruct) DeleteToDoItem(id string) (deleted bool, err error) {
	del, err := ds.inMemoryToDoList.DeleteToDoItem(id)
	return del, err
}
