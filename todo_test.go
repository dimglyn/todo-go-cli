package main

import (
	"testing"
	"math/rand"
	"time"
)

func TestAppendTodo(t *testing.T) {
	repo := TodoRepo{}
	rand.Seed(time.Now().UnixNano())

	id := rand.Intn(50)

	todo := Todo {
		text: "test todo",
		id: id,
	}

	repo, i := AppendTodo(repo, todo)

	if len(repo) != 1{
		t.Error("should add 1 todo in the repo")
	}

	if  i != id + 1 {
		t.Error("should return the next id")
	}
}