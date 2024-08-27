package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	types "github.com/owen-corrigan-bjss/to-do-app/to-do-types"
)

var inMemoryToDoList = types.NewToDoList()
var ids = types.NewCounter()

type ToDoReq struct {
	Description string `json:"description"`
}

type ToDoResponse struct {
	Id          string
	Description string
	Status      bool
}

type CommandType int

const (
	GetCommand = iota
	PostCommand
	UpdateCommand
	DeleteCommand
)

type Request struct {
	reqType     CommandType
	description string
	id          string
	replyChan   chan types.ToDoList
	errorChan   chan error
}

type Server struct {
	requests chan<- Request
}

func NewServer() *Server {
	server := Server{}
	server.requests = startDataManager()
	http.HandleFunc("POST /create", server.HandleCreateNewToDo)
	http.HandleFunc("GET /todos", server.HandleListToDos)
	http.HandleFunc("PUT /update", server.HandleUpdateToDo)
	http.HandleFunc("DELETE /remove", server.HandleDeleteToDo)
	return &server
}

func (s *Server) HandleCreateNewToDo(res http.ResponseWriter, req *http.Request) {

	var toDo ToDoReq
	json.NewDecoder(req.Body).Decode(&toDo)

	if len(toDo.Description) == 0 {
		http.Error(res, "Invalid Request", 400)
		return
	}

	replyChan := make(chan types.ToDoList)
	errorChan := make(chan error)

	s.requests <- Request{PostCommand, toDo.Description, "", replyChan, errorChan}

	responseBody := <-replyChan
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(201)
	json.NewEncoder(res).Encode(responseBody)
}

func (s *Server) HandleListToDos(res http.ResponseWriter, req *http.Request) {

	replyChan := make(chan types.ToDoList)
	errorChan := make(chan error)

	s.requests <- Request{GetCommand, "", "", replyChan, errorChan}

	responseBody := <-replyChan
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(responseBody)

}

func (s *Server) HandleUpdateToDo(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	id := query.Get("id")

	if len(id) == 0 {
		http.Error(res, "Invalid Request", 400)
		return
	}

	replyChan := make(chan types.ToDoList)
	errorChan := make(chan error)

	s.requests <- Request{UpdateCommand, "", id, replyChan, errorChan}

	select {
	case err := <-errorChan:
		http.Error(res, err.Error(), 400)
	case responseBody := <-replyChan:
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		json.NewEncoder(res).Encode(responseBody)
	}
}

func (s *Server) HandleDeleteToDo(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	id := query.Get("id")

	replyChan := make(chan types.ToDoList)
	errorChan := make(chan error)

	s.requests <- Request{DeleteCommand, "", id, replyChan, errorChan}

	select {
	case err := <-errorChan:
		http.Error(res, err.Error(), 400)
	case <-replyChan:
		responseBody := fmt.Sprintf("todo: %s deleted", id)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		json.NewEncoder(res).Encode(responseBody)
	}

}

func startDataManager() chan<- Request {
	requests := make(chan Request)
	go func() {
		for req := range requests {
			switch req.reqType {
			case GetCommand:
				list := inMemoryToDoList.GetToDoMap()

				req.replyChan <- list

			case PostCommand:
				key := inMemoryToDoList.CreateToDoItem(req.description, ids)

				list := make(map[string]types.ToDo)
				list[key] = types.ToDo{Description: req.description, Completed: false}

				req.replyChan <- list

			case UpdateCommand:
				toDo, err := inMemoryToDoList.UpdateToDoItemStatus(req.id)

				if err != nil {
					req.errorChan <- err
				} else {
					list := make(map[string]types.ToDo)

					list[req.id] = toDo

					req.replyChan <- list
				}

			case DeleteCommand:
				_, err := inMemoryToDoList.DeleteToDoItem(req.id)

				if err != nil {
					req.errorChan <- err
				} else {
					list := make(map[string]types.ToDo)
					list[req.id] = types.ToDo{}
					req.replyChan <- list
				}
			default:
				log.Fatal("unknown command type", req.reqType)
			}
		}
	}()
	return requests
}
