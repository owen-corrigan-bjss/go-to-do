package toDos

type ToDo struct {
	Description string
	Complete    bool
}

func NewToDo() *ToDo {
	return &ToDo{}
}

type ToDoListJson struct {
	ToDos []ToDoJson `json:"toDos"`
}

type ToDoJson struct {
	Description string `json:"description"`
	Complete    string `json:"complete"`
}

var ToDoList []ToDo = []ToDo{
	{Description: "this is 0from an array", Complete: false},
	{Description: "test it", Complete: false},
	{Description: "refactor it", Complete: false},
	{Description: "add more stuff", Complete: false},
	{Description: "write 5 more todos", Complete: false},
	{Description: "create a server", Complete: false},
	{Description: "create endpoints", Complete: false},
	{Description: "integrate it with this", Complete: false},
	{Description: "do some more stuff", Complete: false},
	{Description: "write more tests", Complete: false},
}
