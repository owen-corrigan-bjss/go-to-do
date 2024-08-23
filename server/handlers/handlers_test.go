package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddToDoHandler(t *testing.T) {
	t.Run("If the method is not post will return 400", func(t *testing.T) {
		validData := []byte(`{description: "a test task"}`)
		req, err := http.NewRequest("GET", "/todo", bytes.NewBuffer(validData))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleCreateNewToDo(got, req)
		want := 400
		if got.Code != want {
			t.Errorf("wanted %d got %v", want, got.Code)
		}
	})
	t.Run("Creates a new todo and returns the newly created item", func(t *testing.T) {
		validData := []byte(`{description: "a test task"}`)
		req, err := http.NewRequest("POST", "/todo", bytes.NewBuffer(validData))
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
		validData := []byte(``)
		req, err := http.NewRequest("POST", "/todo", bytes.NewBuffer(validData))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		got := httptest.NewRecorder()
		HandleCreateNewToDo(got, req)
		want := 400
		if got.Code != want {
			t.Errorf("wanted %d got %v", want, got.Code)
		}
	})
}
