package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestAppendTodo(t *testing.T) {
	repo := TodoRepo{}
	rand.Seed(time.Now().UnixNano())

	id := rand.Intn(50)

	todo := Todo{
		text: "test todo",
		id:   id,
	}

	repo, i := AppendTodo(repo, todo)

	if len(repo) != 1 {
		t.Error("should add 1 todo in the repo")
	}

	if i != id+1 {
		t.Error("should return the next id")
	}
}

func TestEditTodo(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	id := rand.Intn(50)

	todo := Todo{
		text: "test todo",
		id:   id,
	}

	repo := TodoRepo{todo}
	newText := "new todo"

	_, repo, updatedTodo := EditTodoText(repo, todo.id, newText)

	if updatedTodo.text != newText {
		t.Error("Should update a todo by its id")
	}
}

func TestRemoveTodo(t *testing.T) {
	repo := TodoRepo{
		Todo{
			text: "removable todo",
			id:   4,
		},
	}
	_, repo = RemoveTodo(repo, 4)

	if len(repo) != 0 {
		t.Error("Should remove a todo by it's id")
	}
}

func TestToggleDone(t *testing.T) {
	repo := TodoRepo{
		Todo{
			text: "removable todo",
			id:   4,
			done: false,
		},
	}
	_, repo, todo := ToggleDone(repo, 4)

	if !todo.done {
		t.Error("Should be true (done)")
	}

}
