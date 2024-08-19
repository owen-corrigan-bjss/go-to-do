package main

import (
	"log"
	"os"
	"testing"
	toDos "to-do-app"
	"to-do-app/helpers"
)

func TestMainFunctions(t *testing.T) {
	t.Run("printJsonToDos", func(t *testing.T) {
		toDoJson, err := os.Open("/Users/owen.corrigan/projects/go-to-do/toDos.json")
		if err != nil {
			log.Fatal(err)
		}
		defer toDoJson.Close()

		decodedJson := helpers.DecodeJson(toDoJson)

		got := PrintJsonToDos(decodedJson...)
		want := "Here are your todo's:\n1: This is from the JSON file\n2: test it\n3: refactor it\n4: add more stuff\n5: write 5 more todos\n6: create a server\n7: create endpoints\n8: integrate it with this\n9: do some more stuff\n10: write more tests\n"

		if got != want {
			t.Errorf("wanted %s got %s", want, got)
		}
	})
	t.Run("printToDos", func(t *testing.T) {
		toDoList := []toDos.ToDo{{Description: "a thing", Complete: false}, {Description: "a second thing", Complete: false}}
		got := PrintToDosList(toDoList...)
		want := "Here are your todo's:\n1: a thing\n2: a second thing\n"

		if got != want {
			t.Errorf("wanted %s got %s", want, got)
		}
	})
}
