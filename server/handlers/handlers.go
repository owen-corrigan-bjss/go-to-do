package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	types "to-do-app/to-do-types"
)

var inMemoryToDoList = types.NewToDoList()
var ids = types.NewCounter()

type ToDoReq struct {
	Description string `json:"description"`
}

type ToDoCreatedResponse struct {
	Id          string
	Description string
	Status      bool
}

func HandleCreateNewToDo(res http.ResponseWriter, req *http.Request) {
	var toDo ToDoReq

	json.NewDecoder(req.Body).Decode(&toDo)

	if len(toDo.Description) == 0 {
		fmt.Println("in here")
		http.Error(res, "Invalid Request", 400)
	}

	newItemKey, err := inMemoryToDoList.CreateToDoItem(toDo.Description, ids)

	if err != nil {
		http.Error(res, "Invalid Request", 500)
	}

	newToDo := inMemoryToDoList.GetSingleToDo(newItemKey)

	resBody := ToDoCreatedResponse{newItemKey, newToDo.Description, newToDo.Completed}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(201)
	json.NewEncoder(res).Encode(resBody)
}

func HandleListToDos(res http.ResponseWriter, req *http.Request) {
	toDos := inMemoryToDoList.GetToDoMap()

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(toDos)
}

func Handlers() {
	http.HandleFunc("POST /create", HandleCreateNewToDo)
	http.HandleFunc("GET /todos", HandleListToDos)
}
