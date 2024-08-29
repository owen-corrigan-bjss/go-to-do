package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	types "github.com/owen-corrigan-bjss/to-do-app/to-do-types"
	"github.com/owen-corrigan-bjss/to-do-app/web-app/dataService"
)

type ToDoReq struct {
	Description string `json:"description"`
}

type ToDoResponse struct {
	Id          string
	Description string
	Status      bool
}

type Server struct {
	requests chan<- dataService.Request
}

func NewServer() *Server {
	server := Server{}
	server.requests = dataService.StartDataManager()
	http.HandleFunc("POST /create", server.HandleCreateNewToDo)
	http.HandleFunc("GET /todo-list", server.HandleListToDos)
	http.HandleFunc("GET /todo", server.HandleGetSingleToDo)
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

	defer close(replyChan)
	defer close(errorChan)

	s.requests <- dataService.Request{ReqType: dataService.PostCommand, Description: toDo.Description, Id: "", ReplyChan: replyChan, ErrorChan: errorChan}

	responseBody := <-replyChan
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(201)
	json.NewEncoder(res).Encode(responseBody)
}

func (s *Server) HandleListToDos(res http.ResponseWriter, req *http.Request) {

	replyChan := make(chan types.ToDoList)
	errorChan := make(chan error)
	defer close(replyChan)
	defer close(errorChan)

	s.requests <- dataService.Request{ReqType: dataService.GetCommand, Description: "", Id: "", ReplyChan: replyChan, ErrorChan: errorChan}

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
	defer close(replyChan)
	defer close(errorChan)

	s.requests <- dataService.Request{ReqType: dataService.UpdateCommand, Description: "", Id: id, ReplyChan: replyChan, ErrorChan: errorChan}

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
	defer close(replyChan)
	defer close(errorChan)

	s.requests <- dataService.Request{ReqType: dataService.DeleteCommand, Description: "", Id: id, ReplyChan: replyChan, ErrorChan: errorChan}

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

func (s *Server) HandleGetSingleToDo(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	id := query.Get("id")

	replyChan := make(chan types.ToDoList)
	errorChan := make(chan error)
	defer close(replyChan)
	defer close(errorChan)

	s.requests <- dataService.Request{ReqType: dataService.GetSingleCommand, Description: "", Id: id, ReplyChan: replyChan, ErrorChan: errorChan}

	select {
	case err := <-errorChan:
		http.Error(res, err.Error(), 404)
	case responseBody := <-replyChan:
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		json.NewEncoder(res).Encode(responseBody)
	}
}
