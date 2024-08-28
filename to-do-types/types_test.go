package types

import (
	"testing"
)

func TestCreateToDoItem(t *testing.T) {
	t.Run("will add a single item to the map when there is nothing in the list", func(t *testing.T) {
		toDos := NewToDoList()
		ids := NewCounter()
		toDos.CreateToDoItem("this is item 1", ids)

		got := len(toDos.list)
		want := 1

		if got != want {
			t.Errorf("wanted %d but got %d", want, got)
		}
	})

	t.Run("will add multiple items to the map", func(t *testing.T) {
		toDos := NewToDoList()
		ids := NewCounter()
		toDos.CreateToDoItem("this is item 1", ids)
		toDos.CreateToDoItem("this is item 2", ids)

		got := len(toDos.list)
		want := 2

		if got != want {
			t.Errorf("wanted %d but got %d", want, got)
		}
	})

	t.Run("it will keep creating if you delete some items", func(t *testing.T) {
		toDos := NewToDoList()
		ids := NewCounter()
		toDos.CreateToDoItem("this is item 1", ids)
		toDos.CreateToDoItem("this is item 2", ids)
		toDos.CreateToDoItem("this is item 3", ids)

		toDos.DeleteToDoItem("2")

		toDos.CreateToDoItem("this is a 4th item", ids)

		got := len(toDos.list)
		want := 3

		if got != want {
			t.Errorf("wanted %d but got %d", want, got)
		}
	})
}

func TestGetSingleToDoItem(t *testing.T) {
	t.Run("will return an error if the item doesn't exist", func(t *testing.T) {
		toDos := NewToDoList()

		_, err := toDos.GetSingleToDo("1")
		want := "to do doesn't exist"

		if err.Error() != want {
			t.Errorf("wanted %s but got %s", want, err.Error())
		}
	})

	t.Run("will return a single to do if it exists", func(t *testing.T) {
		toDos := NewToDoList()
		ids := NewCounter()
		toDos.CreateToDoItem("this is item 1", ids)
		toDos.CreateToDoItem("this is item 2", ids)

		got, _ := toDos.GetSingleToDo("1")
		want := "this is item 2"

		if got.Description != want {
			t.Errorf("wanted %s but got %s", want, got.Description)
		}
	})
}

func TestGetToDoMap(t *testing.T) {
	t.Run("will return an empty list if theres nothing in it", func(t *testing.T) {
		toDos := NewToDoList()

		get := toDos.GetToDoMap()

		if len(get) != 0 {
			t.Errorf("wanted %d but got %d", len(get), 0)
		}
	})

	t.Run("will return a single to do if it exists", func(t *testing.T) {
		toDos := NewToDoList()
		ids := NewCounter()
		toDos.CreateToDoItem("this is item 1", ids)
		toDos.CreateToDoItem("this is item 2", ids)

		got := toDos.GetToDoMap()

		if len(got) != 2 {
			t.Errorf("wanted %d but got %d", len(got), 2)
		}
	})
}

func TestUpdateToDoItemStatus(t *testing.T) {
	t.Run("will update an item in the list to be complete", func(t *testing.T) {
		toDos := NewToDoList()
		toDos.list["1"] = ToDo{"a to do", false}
		toDos.UpdateToDoItemStatus("1")

		got := toDos.list["1"].Completed
		want := true

		if got != want {
			t.Errorf("wanted %t but got %t", want, got)
		}
	})

	t.Run("will update an item in the list to not be complete", func(t *testing.T) {
		toDos := NewToDoList()
		toDos.list["1"] = ToDo{"a to do", false}
		toDos.UpdateToDoItemStatus("1")

		got := toDos.list["1"].Completed
		want := true

		if got != want {
			t.Errorf("wanted %t but got %t", want, got)
		}
		toDos.UpdateToDoItemStatus("1")

		got = toDos.list["1"].Completed
		want = false

		if got != want {
			t.Errorf("wanted %t but got %t", want, got)
		}
	})

	t.Run("will return an error if the item doesn't exist", func(t *testing.T) {
		toDos := NewToDoList()
		toDos.list["1"] = ToDo{"this is item 1", false}
		toDos.list["2"] = ToDo{"this is item 2", false}

		_, err := toDos.UpdateToDoItemStatus("3")

		if err == nil {
			t.Errorf("wanted error but got nil")
		}
	})

	t.Run("will return updated item if it exists", func(t *testing.T) {
		toDos := NewToDoList()
		toDos.list["1"] = ToDo{"this is item 1", false}
		toDos.list["2"] = ToDo{"this is item 2", false}

		toDo, _ := toDos.UpdateToDoItemStatus("2")

		if !toDo.Completed {
			t.Errorf("wanted an %t but got %t", true, toDo.Completed)
		}
	})

}

func TestDeleteToDoItemDesc(t *testing.T) {
	t.Run("will delete a give key from a map", func(t *testing.T) {
		toDos := NewToDoList()
		toDos.list["1"] = ToDo{"this is item 1", false}
		toDos.list["2"] = ToDo{"this is item 2", false}
		toDos.list["3"] = ToDo{"this is item 3", false}
		toDos.DeleteToDoItem("2")

		_, ok := toDos.list["2"]

		if ok != false {
			t.Errorf("wanted %t but got %t", false, ok)
		}
	})
	t.Run("will return an error if the item doesn't exist", func(t *testing.T) {
		toDos := NewToDoList()
		toDos.list["1"] = ToDo{"this is item 1", false}
		toDos.list["2"] = ToDo{"this is item 2", false}

		_, err := toDos.DeleteToDoItem("3")

		if err == nil {
			t.Errorf("wanted error but got nil")
		}
	})

	t.Run("will return nil if the item exists", func(t *testing.T) {
		toDos := NewToDoList()
		toDos.list["1"] = ToDo{"this is item 1", false}
		toDos.list["2"] = ToDo{"this is item 2", false}

		_, err := toDos.DeleteToDoItem("2")

		if err != nil {
			t.Errorf("wanted an nil but got %v", err)
		}
	})

}
func TestIncrementCounter(t *testing.T) {
	t.Run("increaments the count by 1", func(t *testing.T) {
		newCount := IdCounter{}
		newCount.count = 0

		newCount.IncrementCounter()
		got := newCount.count
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("increaments multiple times", func(t *testing.T) {
		newCount := IdCounter{}
		newCount.count = 0

		newCount.IncrementCounter()
		newCount.IncrementCounter()
		newCount.IncrementCounter()
		newCount.IncrementCounter()
		got := newCount.count
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestGetNewId(t *testing.T) {
	t.Run("gets the ID and increaments the count", func(t *testing.T) {
		newCount := IdCounter{}
		newCount.count = 0

		got := newCount.GetNewId()
		want := "0"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		if newCount.count != 1 {
			t.Errorf("got %d want %d", newCount.count, 1)
		}
	})
}
