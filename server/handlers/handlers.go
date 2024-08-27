package handlers

import (
	"encoding/json"
	"fmt"
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

func HandleCreateNewToDo(res http.ResponseWriter, req *http.Request) {
	var toDo ToDoReq

	json.NewDecoder(req.Body).Decode(&toDo)

	if len(toDo.Description) == 0 {
		http.Error(res, "Invalid Request", 400)
		return
	}
	newItemKey := inMemoryToDoList.CreateToDoItem(toDo.Description, ids)

	resBody := ToDoResponse{newItemKey, toDo.Description, false}

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

func HandleUpdateToDo(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	id := query.Get("id")

	err := inMemoryToDoList.UpdateToDoItemStatus(id)

	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	updatedToDo := inMemoryToDoList.GetSingleToDo(id)
	responseBody := ToDoResponse{id, updatedToDo.Description, updatedToDo.Completed}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(responseBody)
}

func HandleDeleteToDo(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	id := query.Get("id")

	err := inMemoryToDoList.DeleteToDoItem(id)

	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	responseBody := fmt.Sprintf("todo: %s deleted", id)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(responseBody)

}

func Handlers() {
	http.HandleFunc("POST /create", HandleCreateNewToDo)
	http.HandleFunc("GET /todos", HandleListToDos)
	http.HandleFunc("PUT /update", HandleUpdateToDo)
	http.HandleFunc("DELETE /remove", HandleDeleteToDo)
}
