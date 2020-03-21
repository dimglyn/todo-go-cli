package main

import (
	"fmt"
)

type Todo struct {
	text string
	id   int
}

type TodoRepo []Todo

func RenderTodos(tr TodoRepo) {
	fmt.Println(tr)
}

func AppendTodo(tr TodoRepo, todo Todo) (TodoRepo, int) {
	tr = append(tr, todo)
	return tr, todo.id + 1
}
