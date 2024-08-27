package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	types "github.com/owen-corrigan-bjss/to-do-app/to-do-types"
)

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Status Incorrect, got %d, want %d", got, want)
	}
}

func TestHandleCreateNewToDo(t *testing.T) {
	t.Run("Creates a new todo and returns the newly created item", func(t *testing.T) {
		validData := []byte(`{"description": "a test task"}`)
		req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(validData))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleCreateNewToDo(got, req)

		assertStatus(t, got.Code, 201)

		var responseBody ToDoResponse
		json.NewDecoder(got.Body).Decode(&responseBody)

		wantBody := ToDoResponse{"0", "a test task", false}

		if !reflect.DeepEqual(responseBody, wantBody) {
			t.Errorf("wanted %v got %v", wantBody, got.Body.String())
		}

	})
	t.Run("if there is no description will return a 400", func(t *testing.T) {
		validData := []byte(`{"description": ""}`)
		req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(validData))
		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleCreateNewToDo(got, req)

		assertStatus(t, got.Code, 400)
	})
}

func TestHandleUpdateToDo(t *testing.T) {
	t.Run("will update a todo", func(t *testing.T) {
		inMemoryToDoList = types.NewToDoList()
		ids = types.NewCounter()
		inMemoryToDoList.CreateToDoItem("test to do", ids)
		req, _ := http.NewRequest("PUT", "/update?id=0", nil)
		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleUpdateToDo(got, req)

		assertStatus(t, got.Code, 200)

		updatedItemStatus := inMemoryToDoList.GetSingleToDo("0").Completed

		if updatedItemStatus != true {
			t.Errorf("wanted %t for %t", true, updatedItemStatus)
		}
	})
	t.Run("will return an error if the item doesn't exist", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/update?id=10", nil)
		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleUpdateToDo(got, req)

		assertStatus(t, got.Code, 400)
	})
}

func TestHandleDeleteToDo(t *testing.T) {
	t.Run("will delete a todo", func(t *testing.T) {
		inMemoryToDoList = types.NewToDoList()
		ids = types.NewCounter()
		inMemoryToDoList.CreateToDoItem("test to do", ids)
		req, _ := http.NewRequest("DELETE", "/remove?id=0", nil)
		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleDeleteToDo(got, req)

		assertStatus(t, got.Code, 200)
	})
	t.Run("will return an error if the item doesn't exist", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/update?id=10", nil)
		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleDeleteToDo(got, req)

		assertStatus(t, got.Code, 400)
	})
}

func TestHandleListToDos(t *testing.T) {
	t.Run("If there is nothing in the list returns an empty list", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/todos", nil)

		request.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleListToDos(got, request)

		assertStatus(t, got.Code, 200)

		var responseBody ToDoResponse
		json.NewDecoder(got.Body).Decode(&responseBody)

		wantBody := ToDoResponse{}

		if !reflect.DeepEqual(responseBody, wantBody) {
			t.Errorf("wanted %v got %v", wantBody, got.Body.String())
		}
	})

	t.Run("returns the list", func(t *testing.T) {
		inMemoryToDoList = types.NewToDoList()
		ids = types.NewCounter()
		inMemoryToDoList.CreateToDoItem("heres a to do", ids)
		req, _ := http.NewRequest("GET", "/todo", nil)

		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleListToDos(got, req)

		assertStatus(t, got.Code, 200)

		var responseBody types.ToDoList
		json.NewDecoder(got.Body).Decode(&responseBody)

		wantBody := types.ToDo{Description: "heres a to do", Completed: false}

		if !reflect.DeepEqual(responseBody["0"], wantBody) {
			t.Errorf("wanted %v got %v", wantBody, got.Body.String())
		}
	})
}
