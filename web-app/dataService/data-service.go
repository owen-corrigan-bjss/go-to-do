package dataService

import (
	"log"

	types "github.com/owen-corrigan-bjss/to-do-app/to-do-types"
)

var ds *DataStoreStruct = NewDataStore()

type CommandType int

const (
	GetCommand = iota
	PostCommand
	UpdateCommand
	DeleteCommand
	GetSingleCommand
)

type Request struct {
	ReqType     CommandType
	Description string
	Id          string
	ReplyChan   chan types.ToDoList
	ErrorChan   chan error
}

func StartDataManager() chan<- Request {
	requests := make(chan Request)
	go func() {
		for req := range requests {
			switch req.ReqType {
			case GetCommand:
				list := ds.GetToDoMap()

				req.ReplyChan <- list

			case PostCommand:
				key := ds.CreateToDoItem(req.Description)

				list := make(map[string]types.ToDo)
				list[key] = types.ToDo{Description: req.Description, Completed: false}

				req.ReplyChan <- list

			case UpdateCommand:
				toDo, err := ds.UpdateToDoItemStatus(req.Id)

				if err != nil {
					req.ErrorChan <- err
				} else {
					list := make(map[string]types.ToDo)

					list[req.Id] = toDo

					req.ReplyChan <- list
				}

			case DeleteCommand:
				_, err := ds.DeleteToDoItem(req.Id)

				if err != nil {
					req.ErrorChan <- err
				} else {
					list := make(map[string]types.ToDo)
					list[req.Id] = types.ToDo{}
					req.ReplyChan <- list
				}
			case GetSingleCommand:
				toDo, err := ds.GetSingleToDo(req.Id)

				if err != nil {
					req.ErrorChan <- err
				} else {
					list := make(map[string]types.ToDo)

					list[req.Id] = toDo

					req.ReplyChan <- list
				}

			default:
				log.Fatal("unknown command type", req.ReqType)
			}
		}
	}()
	return requests
}
