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

	_, repo, updatedTodo := EditTodoByID(repo, todo.id, newText)

	if updatedTodo.text != newText {
		t.Error("Should update a todo by its id")
	}
}
