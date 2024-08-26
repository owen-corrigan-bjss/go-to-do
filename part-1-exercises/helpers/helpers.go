package helpers

import (
	"encoding/json"
	toDos "github.com/owen-corrigan-bjss/to-do-app/part-1-exercises"
	"io"
	"os"
)

func DecodeJson(osFile *os.File) []toDos.ToDoJson {
	byteArr, _ := io.ReadAll(osFile)
	var toDoObject toDos.ToDoListJson
	json.Unmarshal(byteArr, &toDoObject)
	return toDoObject.ToDos
}
