package main

import (
	toDos "github.com/owen-corrigan-bjss/to-do-app/part-1-exercises"
	"testing"
)

func BenchmarkReadToDoDesc(b *testing.B) {
	var list LockedToDo
	list.toDos = []toDos.ToDo{{Description: "a thing", Complete: false}, {Description: "a second thing", Complete: false}}
	done := make(chan bool)

	for range b.N {
		go ReadToDoDesc(&list, 1, done)
		<-done
	}
}
