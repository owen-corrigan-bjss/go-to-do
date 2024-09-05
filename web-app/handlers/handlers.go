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

type Handlers struct {
	requests chan<- dataService.Request
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "content-type")
}

func NewHandlers() *Handlers {
	handlers := Handlers{}
	handlers.requests = dataService.StartDataManager()
	http.HandleFunc("POST /create", handlers.HandleCreateNewToDo)
	http.HandleFunc("OPTIONS /create", handlers.HandleOptions)
	http.HandleFunc("GET /todo-list", handlers.HandleListToDos)
	http.HandleFunc("GET /todo", handlers.HandleGetSingleToDo)
	http.HandleFunc("PUT /update", handlers.HandleUpdateToDo)
	http.HandleFunc("OPTIONS /update", handlers.HandleOptions)
	http.HandleFunc("DELETE /remove", handlers.HandleDeleteToDo)
	http.HandleFunc("OPTIONS /remove", handlers.HandleOptions)
	return &handlers
}

func (s *Handlers) HandleOptions(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
}

func (s *Handlers) HandleCreateNewToDo(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
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

func (s *Handlers) HandleListToDos(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
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

func (s *Handlers) HandleUpdateToDo(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
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

func (s *Handlers) HandleDeleteToDo(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
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

func (s *Handlers) HandleGetSingleToDo(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
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
