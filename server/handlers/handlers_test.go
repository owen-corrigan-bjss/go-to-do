package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	types "github.com/owen-corrigan-bjss/to-do-app/to-do-types"
)

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
		want := 201
		if got.Code != want {
			t.Errorf("wanted %d got %v", want, got.Code)
		}
	})
	t.Run("if there is no description will return a 400", func(t *testing.T) {
		validData := []byte(`{"description": ""}`)
		req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(validData))
		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleCreateNewToDo(got, req)
		want := 400
		if got.Code != want {
			t.Errorf("wanted %d got %v", want, got.Code)
		}
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

		if got.Code != 200 {
			t.Errorf("wanted %d for %d", 200, got.Code)
		}

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

		if got.Code != 400 {
			t.Errorf("wanted %d for %d", 400, got.Code)
		}

		// var bodyString string
		// json.Unmarshal(got.Body.Bytes(), &bodyString)

		// if bodyString != "item doesn't exist" {
		// 	t.Errorf("wanted %s got %s", got.Body.String(), "item doesn't exist")
		// }
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

		if got.Code != 200 {
			t.Errorf("wanted %d for %d", 200, got.Code)
		}
	})
	t.Run("will return an error if the item doesn't exist", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/update?id=10", nil)
		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleDeleteToDo(got, req)

		if got.Code != 400 {
			t.Errorf("wanted %d for %d", 400, got.Code)
		}

		// if got.Body.String() != "item doesn't exist" {
		// 	t.Errorf("wanted %s got %s", got.Body.String(), "item doesn't exist")
		// }
	})
}

// func TestHandleListToDos(t *testing.T) {
// 	t.Run("If there is nothing in the list returns an empty list", func(t *testing.T) {
// 		request, _ := http.NewRequest("GET", "/todos", nil)

// 		request.Header.Set("Content-Type", "application/json")
// 		response := httptest.NewRecorder()
// 		HandleListToDos(response, request)
// 		want := 200
// 		if response.Code != want {
// 			t.Errorf("wanted %d got %v", want, response.Code)
// 		}
// 		//comparison not working
// 		if response.Body.String() != "{}" {
// 			t.Errorf("wanted %s got %s", "{}", response.Body.String())
// 		}
// 	})

// 	t.Run("returns the list", func(t *testing.T) {

// 		inMemoryToDoList.CreateToDoItem("heres a to do", ids)
// 		req, err := http.NewRequest("GET", "/todo", nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		req.Header.Set("Content-Type", "application/json")
// 		got := httptest.NewRecorder()
// 		HandleListToDos(got, req)
// 		want := 200
// 		if got.Code != want {
// 			t.Errorf("wanted %d got %v", want, got.Code)
// 		}
// 		//fix this
// 		wantString := "{\"0\":{\"Description\":\"heres a to do\",\"Completed\":false}}"
// 		if got.Body.String() != wantString {
// 			t.Errorf("wanted %s got %s", wantString, got.Body.String())
// 		}
// 	})
// }
