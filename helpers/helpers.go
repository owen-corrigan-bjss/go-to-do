package helpers

import (
	"encoding/json"
	"io"
	"os"
	toDos "to-do-app"
)

func DecodeJson(osFile *os.File) []toDos.ToDoJson {
	byteArr, _ := io.ReadAll(osFile)
	var toDoObject toDos.ToDoListJson
	json.Unmarshal(byteArr, &toDoObject)
	return toDoObject.ToDos
}
